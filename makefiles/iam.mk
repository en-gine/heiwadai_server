# IAM management for GitHub Actions
AWS_REGION ?= ap-northeast-1
AWS_PROFILE ?= heiwadai
AWS_CMD = aws --profile $(AWS_PROFILE)
GITHUB_ACTIONS_USER = github-actions-ecr-user
ECR_POLICY_NAME = GitHubActionsECRPolicy

# Create GitHub Actions IAM user for ECR access
create-github-actions-user:
	@echo "Creating IAM user for GitHub Actions..."
	$(AWS_CMD) iam create-user --user-name $(GITHUB_ACTIONS_USER) || true
	@echo "Creating ECR access policy..."
	@echo '{ \
		"Version": "2012-10-17", \
		"Statement": [ \
			{ \
				"Effect": "Allow", \
				"Action": [ \
					"ecr:GetAuthorizationToken" \
				], \
				"Resource": "*" \
			}, \
			{ \
				"Effect": "Allow", \
				"Action": [ \
					"ecr:BatchCheckLayerAvailability", \
					"ecr:GetDownloadUrlForLayer", \
					"ecr:BatchGetImage", \
					"ecr:BatchImportLayerPart", \
					"ecr:InitiateLayerUpload", \
					"ecr:UploadLayerPart", \
					"ecr:CompleteLayerUpload", \
					"ecr:PutImage", \
					"ecr:DescribeRepositories", \
					"ecr:DescribeImages" \
				], \
				"Resource": "arn:aws:ecr:$(AWS_REGION):$(AWS_ID):repository/heiwadai-server" \
			} \
		] \
	}' > /tmp/github-actions-ecr-policy.json
	$(AWS_CMD) iam create-policy \
		--policy-name $(ECR_POLICY_NAME) \
		--policy-document file:///tmp/github-actions-ecr-policy.json \
		--description "Policy for GitHub Actions to access heiwadai-server ECR repository" || true
	@echo "Attaching policy to user..."
	$(AWS_CMD) iam attach-user-policy \
		--user-name $(GITHUB_ACTIONS_USER) \
		--policy-arn arn:aws:iam::$(AWS_ID):policy/$(ECR_POLICY_NAME)
	@rm -f /tmp/github-actions-ecr-policy.json

# Create access key for GitHub Actions user
create-github-actions-access-key:
	@echo "Creating access key for GitHub Actions user..."
	@echo "IMPORTANT: Save these credentials securely!"
	@echo "=========================================="
	$(AWS_CMD) iam create-access-key --user-name $(GITHUB_ACTIONS_USER)
	@echo "=========================================="
	@echo "Add these to GitHub Secrets:"
	@echo "- AWS_ACCESS_KEY_ID: <AccessKeyId from above>"
	@echo "- AWS_SECRET_ACCESS_KEY: <SecretAccessKey from above>"

# Delete GitHub Actions access key
delete-github-actions-access-key:
ifndef ACCESS_KEY_ID
	$(error ACCESS_KEY_ID is required. Usage: make delete-github-actions-access-key ACCESS_KEY_ID=AKIA...)
endif
	@echo "Deleting access key $(ACCESS_KEY_ID)..."
	$(AWS_CMD) iam delete-access-key \
		--user-name $(GITHUB_ACTIONS_USER) \
		--access-key-id $(ACCESS_KEY_ID)

# List access keys for GitHub Actions user
list-github-actions-access-keys:
	@echo "Access keys for $(GITHUB_ACTIONS_USER):"
	$(AWS_CMD) iam list-access-keys --user-name $(GITHUB_ACTIONS_USER)

# Test GitHub Actions credentials
test-github-actions-credentials:
ifndef AWS_ACCESS_KEY_ID
	$(error AWS_ACCESS_KEY_ID is required for testing)
endif
ifndef AWS_SECRET_ACCESS_KEY
	$(error AWS_SECRET_ACCESS_KEY is required for testing)
endif
	@echo "Testing GitHub Actions credentials..."
	@echo "Testing ECR repository access..."
	@AWS_ACCESS_KEY_ID=$(AWS_ACCESS_KEY_ID) \
	 AWS_SECRET_ACCESS_KEY=$(AWS_SECRET_ACCESS_KEY) \
	 AWS_DEFAULT_REGION=$(AWS_REGION) \
	 aws ecr describe-repositories --repository-names heiwadai-server
	@echo "Testing ECR login..."
	@AWS_ACCESS_KEY_ID=$(AWS_ACCESS_KEY_ID) \
	 AWS_SECRET_ACCESS_KEY=$(AWS_SECRET_ACCESS_KEY) \
	 AWS_DEFAULT_REGION=$(AWS_REGION) \
	 aws ecr get-login-password --region $(AWS_REGION) > /dev/null && echo "✅ ECR login successful"
	@echo "✅ All tests passed! Credentials are working correctly."

# Delete GitHub Actions user and associated resources
delete-github-actions-user:
	@echo "WARNING: This will delete the GitHub Actions user and all associated access keys!"
	@read -p "Are you sure? (y/N): " confirm && [ "$$confirm" = "y" ]
	@echo "Detaching policies..."
	$(AWS_CMD) iam detach-user-policy \
		--user-name $(GITHUB_ACTIONS_USER) \
		--policy-arn arn:aws:iam::$(AWS_ID):policy/$(ECR_POLICY_NAME) || true
	@echo "Deleting access keys..."
	@for key in $$($(AWS_CMD) iam list-access-keys --user-name $(GITHUB_ACTIONS_USER) --query 'AccessKeyMetadata[].AccessKeyId' --output text 2>/dev/null || true); do \
		echo "Deleting access key: $$key"; \
		$(AWS_CMD) iam delete-access-key --user-name $(GITHUB_ACTIONS_USER) --access-key-id $$key || true; \
	done
	@echo "Deleting user..."
	$(AWS_CMD) iam delete-user --user-name $(GITHUB_ACTIONS_USER) || true
	@echo "Deleting policy..."
	$(AWS_CMD) iam delete-policy --policy-arn arn:aws:iam::$(AWS_ID):policy/$(ECR_POLICY_NAME) || true
	@echo "✅ GitHub Actions user and resources deleted."

# Show current GitHub Actions user status
show-github-actions-user:
	@echo "GitHub Actions User Status:"
	@echo "=========================="
	@echo "User: $(GITHUB_ACTIONS_USER)"
	@$(AWS_CMD) iam get-user --user-name $(GITHUB_ACTIONS_USER) || echo "❌ User not found"
	@echo ""
	@echo "Attached Policies:"
	@$(AWS_CMD) iam list-attached-user-policies --user-name $(GITHUB_ACTIONS_USER) || echo "❌ No policies found"
	@echo ""
	@echo "Access Keys:"
	@$(AWS_CMD) iam list-access-keys --user-name $(GITHUB_ACTIONS_USER) || echo "❌ No access keys found"

# Help for IAM commands
help-iam:
	@echo "IAM Management Commands for GitHub Actions:"
	@echo "========================================="
	@echo "  create-github-actions-user     - Create IAM user and ECR policy"
	@echo "  create-github-actions-access-key - Create access key for GitHub Actions"
	@echo "  delete-github-actions-access-key ACCESS_KEY_ID=... - Delete specific access key"
	@echo "  list-github-actions-access-keys - List all access keys for the user"
	@echo "  test-github-actions-credentials AWS_ACCESS_KEY_ID=... AWS_SECRET_ACCESS_KEY=... - Test credentials"
	@echo "  show-github-actions-user       - Show current user status"
	@echo "  delete-github-actions-user     - Delete user and all resources"
	@echo ""
	@echo "Example workflow:"
	@echo "  1. make create-github-actions-user"
	@echo "  2. make create-github-actions-access-key"
	@echo "  3. Add credentials to GitHub Secrets"
	@echo "  4. make test-github-actions-credentials AWS_ACCESS_KEY_ID=... AWS_SECRET_ACCESS_KEY=..."

.PHONY: create-github-actions-user create-github-actions-access-key delete-github-actions-access-key \
        list-github-actions-access-keys test-github-actions-credentials delete-github-actions-user \
        show-github-actions-user help-iam