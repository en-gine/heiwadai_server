// 固定IPプロキシLambda関数
const axios = require('axios');

exports.handler = async (event) => {
    const { 
        url, 
        method = 'GET', 
        headers = {}, 
        data = null 
    } = JSON.parse(event.body || '{}');
    
    try {
        const response = await axios({
            url,
            method,
            headers,
            data,
            timeout: 30000
        });
        
        return {
            statusCode: 200,
            body: JSON.stringify({
                status: response.status,
                data: response.data
            })
        };
    } catch (error) {
        return {
            statusCode: error.response?.status || 500,
            body: JSON.stringify({
                error: error.message,
                details: error.response?.data
            })
        };
    }
};