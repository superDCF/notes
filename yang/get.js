const querystring = require("querystring")
const http = require("http")
const fs = require("fs")
// fs.writeFileSync('./data.json', `{"id":13514863,"contacter":"陈实果然"}`, 'utf8', ress=>{
//     console.log("ress",ress)
// });

const options = {
    method: 'POST',
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        Cookie: 'OUTFOX_SEARCH_USER_ID_NCOO=365594159.8071182; uudid=cms5c3bc48c-ea65-2c99-9260-1b9140b334bf; acw_tc=2f624a2316441558841242852e5428c60be9a27859c132c33404260dfe66a7; JSESSIONID=CEF230690E6A94D66B6E4BF5530FAEF1; Mcc=4F13C1F88FE88E0D7CB0BCD31B48A3E7; Identify=AAE1399D1B6D7288B64490A9DD58479955C750A8CB3504A4B5D9335575B0B441573DFFB33C6B3463',
        Host: 'crm.soqi.cn',
        Referer: 'http://crm.soqi.cn/baibao/index.html',
        'User-Agent': 'Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1',
        'Content-Length': 114
    }
};

// 将数据写入请求主体。
const uid = [


]
const notPay = "http://crm.soqi.cn/notPayResources/notPayResourcesList.xhtml"
const newSrc =  "http://crm.soqi.cn/customHome/getNewCustomList.xhtml"
function postUid(uid) {
    options.headers["Content-Length"] = 114;
    const req = http.request(newSrc, options, (res) => {
        console.log(`状态码: ${res.statusCode}`);
        console.log(`响应头: ${JSON.stringify(res.headers)}`);
        res.setEncoding('utf8');
        let str = ''
        res.on('data', (chunk) => {
            str += chunk
            setTimeout(() => {
                let arr = str.match(/(?<="id":)\d+/g)

                fs.writeFileSync('./data.json', JSON.stringify(arr), 'utf8', ress => {
                    console.log("ress", ress)
                });
            }, 500)
            console.log(`响应主体: ${str}`);
        });
        res.on('end', () => {
            console.log('响应中已无数据');
        });
    });

    req.on('error', (e) => {
        console.error(`请求遇到问题: ${e.message}`);
    });
    req.write(`resourceType=0&isContact=0&orderType=0&isOceanLock=0&companyId=0&index=0&num=100&startTime=&endTime=&protectTime=0`);
    req.end();
}
postUid()
// uid.forEach(id => {
//     postUid(id)
// })