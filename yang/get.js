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
        Cookie: 'acw_tc=781bad2815833236511914467e20d3f2881e2650cd8d0828893eb57ddcb6f9; JSESSIONID=CC07E1331F4FB7D32324D19081CE437D; Mcc=671AF6E631C551B77CB0BCD31B48A3E7; Identify=DAD97D9B92783535AEA4CC11E3FF1FB520D0F4E9942DD98600525EDCC42CCBAE3230A38E1626259D',
        Host: 'crm.soqi.cn',
        Referer: 'http://crm.soqi.cn/baibao/index.html',
        'User-Agent': 'Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1'
    }
};

// 将数据写入请求主体。
const uid = [


]
const notPay = "http://crm.soqi.cn/notPayResources/notPayResourcesList.xhtml"
const newSrc =  "http://crm.soqi.cn/customHome/getNewCustomList.xhtml"
function postUid(uid) {
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