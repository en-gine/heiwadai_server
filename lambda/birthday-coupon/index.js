const axios = require('axios');
const https = require('https');

/**
 * Discord/Slack Webhook通知関数（固定IP不要）
 */
async function sendWebhookNotification(subject, content, isError = false) {
    const webhookUrl = process.env.WEBHOOK_URL;
    
    if (!webhookUrl) {
        console.log('Webhook URL not configured. Skipping notification.');
        return;
    }

    const emoji = isError ? '🚨' : '✅';
    const color = isError ? 16711680 : 65280; // Red or Green
    
    // Discord Webhook format
    const payload = {
        embeds: [{
            title: `${emoji} [Heiwadai] ${subject}`,
            description: `\`\`\`\n${content}\n\`\`\``,
            color: color,
            timestamp: new Date().toISOString(),
            footer: {
                text: "Heiwadai Birthday Coupon System"
            }
        }]
    };

    try {
        await makeHttpRequest(webhookUrl, JSON.stringify(payload), {
            'Content-Type': 'application/json'
        });
        console.log('Webhook notification sent successfully');
    } catch (error) {
        console.error('Failed to send webhook notification:', error.message);
    }
}

/**
 * HTTP リクエスト関数
 */
function makeHttpRequest(url, postData, headers) {
    return new Promise((resolve, reject) => {
        const urlObj = new URL(url);
        const options = {
            hostname: urlObj.hostname,
            port: urlObj.port || (urlObj.protocol === 'https:' ? 443 : 80),
            path: urlObj.pathname + urlObj.search,
            method: 'POST',
            headers: {
                'Content-Length': Buffer.byteLength(postData),
                ...headers
            }
        };
        
        const req = https.request(options, (res) => {
            let data = '';
            res.on('data', (chunk) => data += chunk);
            res.on('end', () => {
                if (res.statusCode >= 200 && res.statusCode < 300) {
                    resolve({ statusCode: res.statusCode, body: data });
                } else {
                    reject(new Error(`HTTP ${res.statusCode}: ${data}`));
                }
            });
        });
        
        req.on('error', reject);
        req.write(postData);
        req.end();
    });
}

/**
 * Lambda handler
 */
exports.handler = async (event) => {
    console.log('Starting birthday coupon issuance...');
    
    const startTime = new Date();
    let result = null;
    let error = null;
    
    try {
        const response = await axios.post(
            process.env.CRON_ACCESS_ENDPOINT,
            {},
            {
                headers: {
                    'Authorization': process.env.CRON_ACCESS_SECRET,
                    'X-Cron-Key': process.env.CRON_ACCESS_KEY,
                    'Content-Type': 'application/json',
                },
                timeout: 30000 // 30秒タイムアウト
            }
        );
        
        console.log('Birthday coupon issued successfully:', response.data);
        
        const affectedUserCount = response.data.affectedUserCount || 0;
        const endTime = new Date();
        const duration = endTime - startTime;
        
        result = {
            statusCode: 200,
            body: JSON.stringify({
                message: 'Birthday coupon issued successfully',
                affectedUserCount: affectedUserCount,
                timestamp: new Date().toISOString(),
                duration: `${duration}ms`
            })
        };

        // 成功メール通知
        const successContent = `誕生日クーポンの発行が正常に完了しました。

📊 実行結果:
• 対象ユーザー数: ${affectedUserCount}人
• 実行時間: ${duration}ms
• 開始時刻: ${startTime.toLocaleString('ja-JP', { timeZone: 'Asia/Tokyo' })}
• 完了時刻: ${endTime.toLocaleString('ja-JP', { timeZone: 'Asia/Tokyo' })}

🔗 詳細:
API エンドポイント: ${process.env.CRON_ACCESS_ENDPOINT}
レスポンス: ${JSON.stringify(response.data, null, 2)}`;

        await sendWebhookNotification('誕生日クーポン発行完了', successContent, false);
        
        return result;
        
    } catch (err) {
        error = err;
        const endTime = new Date();
        const duration = endTime - startTime;
        
        console.error('Failed to issue birthday coupon:', error.message);
        
        if (error.response) {
            console.error('Response status:', error.response.status);
            console.error('Response data:', error.response.data);
        }

        result = {
            statusCode: 500,
            body: JSON.stringify({
                error: 'Failed to issue birthday coupon',
                message: error.message,
                timestamp: new Date().toISOString(),
                duration: `${duration}ms`
            })
        };

        // エラーメール通知
        let errorContent = `❌ 誕生日クーポンの発行中にエラーが発生しました。

🚨 エラー詳細:
• エラーメッセージ: ${error.message}
• 実行時間: ${duration}ms
• 開始時刻: ${startTime.toLocaleString('ja-JP', { timeZone: 'Asia/Tokyo' })}
• エラー発生時刻: ${endTime.toLocaleString('ja-JP', { timeZone: 'Asia/Tokyo' })}

🔍 デバッグ情報:
API エンドポイント: ${process.env.CRON_ACCESS_ENDPOINT}`;

        if (error.response) {
            errorContent += `
HTTP ステータス: ${error.response.status}
レスポンス: ${JSON.stringify(error.response.data, null, 2)}`;
        }

        errorContent += `

📋 対応が必要な場合は、AWS CloudWatch Logsでより詳細なログを確認してください。
Log Group: /aws/lambda/birthday-coupon-function`;

        await sendWebhookNotification('誕生日クーポン発行エラー', errorContent, true);
        
        // エラーでもLambdaは成功として扱う（CloudWatchでログ確認可能）
        return result;
    }
};