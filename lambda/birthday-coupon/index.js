const axios = require('axios');
const sgMail = require('@sendgrid/mail');

// SendGrid設定
sgMail.setApiKey(process.env.SENDGRID_API_KEY);

/**
 * メール送信関数
 */
async function sendEmail(subject, content, isError = false) {
    if (!process.env.SENDGRID_API_KEY || !process.env.NOTIFICATION_EMAIL) {
        console.log('SendGrid API key or notification email not configured. Skipping email notification.');
        return;
    }

    const msg = {
        to: process.env.NOTIFICATION_EMAIL,
        from: process.env.MAIL_FROM || 'no-reply@heiwadai-hotel.app',
        subject: `[Heiwadai] ${subject}`,
        text: content,
        html: `
            <div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
                <h2 style="color: ${isError ? '#dc3545' : '#28a745'};">
                    ${isError ? '🚨' : '✅'} ${subject}
                </h2>
                <div style="background-color: #f8f9fa; padding: 20px; border-radius: 5px; margin: 20px 0;">
                    <pre style="white-space: pre-wrap; font-family: monospace; font-size: 14px;">${content}</pre>
                </div>
                <p style="color: #6c757d; font-size: 12px;">
                    送信時刻: ${new Date().toLocaleString('ja-JP', { timeZone: 'Asia/Tokyo' })}
                </p>
            </div>
        `
    };

    try {
        await sgMail.send(msg);
        console.log('Email notification sent successfully');
    } catch (error) {
        console.error('Failed to send email notification:', error.message);
        if (error.response) {
            console.error('SendGrid response:', error.response.body);
        }
    }
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

        await sendEmail('誕生日クーポン発行完了', successContent, false);
        
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
        const errorContent = `❌ 誕生日クーポンの発行中にエラーが発生しました。

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

        await sendEmail('誕生日クーポン発行エラー', errorContent, true);
        
        // エラーでもLambdaは成功として扱う（CloudWatchでログ確認可能）
        return result;
    }
};