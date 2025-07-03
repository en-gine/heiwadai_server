const axios = require('axios');

exports.handler = async (event) => {
    console.log('Starting birthday coupon issuance...');
    
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
        
        return {
            statusCode: 200,
            body: JSON.stringify({
                message: 'Birthday coupon issued successfully',
                affectedUserCount: response.data.affectedUserCount || 0,
                timestamp: new Date().toISOString()
            })
        };
        
    } catch (error) {
        console.error('Failed to issue birthday coupon:', error.message);
        
        if (error.response) {
            console.error('Response status:', error.response.status);
            console.error('Response data:', error.response.data);
        }
        
        // エラーでもLambdaは成功として扱う（CloudWatchでログ確認可能）
        return {
            statusCode: 500,
            body: JSON.stringify({
                error: 'Failed to issue birthday coupon',
                message: error.message,
                timestamp: new Date().toISOString()
            })
        };
    }
};