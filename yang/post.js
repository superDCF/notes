const querystring = require("querystring")
const http = require("http")
const fs = require("fs")
const options = {
    method: 'POST',
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        Cookie: 'acw_tc=76b20f6816213390654208282e05f899f33a036802d585676a9dc1bbfd4c7c; JSESSIONID=59023E8525602FC7441224D8A25B1D94; OUTFOX_SEARCH_USER_ID_NCOO=36728864.31397198; Mcc=9C8D40C460D4FF7B7CB0BCD31B48A3E7; Identify=5B4CF142207A97E73F7621DEE241423567D6598716256962DF49DE38486028F993CC785D709E1471',
        Host: 'crm.soqi.cn',
        Referer: 'http://crm.soqi.cn/baibao/index.html',
        'User-Agent': 'Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1',
        'Content-Length': 114
    }
};

// 将数据写入请求主体。
let data = fs.readFileSync('./data.json', 'utf8')
console.log(JSON.parse(data))
console.log(typeof JSON.parse(data))
console.log(JSON.parse(data)[0])

// const uid = ["13438137","13436295","13436107","13435810","13435776","13435777","13427510","13427123","13427031","13426858","13426561","13426508","13426474","13426453","13426096","13425639","13425272","13424862","13424843","13424055","13423895","13423518","13422230","13422182","13422028","13422022","13421727","13421478","13420841","13420787","13420663","13420307","13420297","13420179","13419853","13419838","13419802","13419712","13419666","13419143","13418197","13417510","13416768","13415790","13415501","13415502","13415440","13415348","13415258","13415081","13414161","13412777","13412743","13412568","13410930","13410855","13410658","13410626","13410594","13410390","13410196","13409875","13408603","13408540","13408402","13408117","13408104","13408013","13407969","13407961","13407931","13407882","13407793","13407740","13407705","13406865","13406754","13406589","13406542","13406153","13405395","13404845","13404664","13404117","13403776","13402351","13401704","13400944","13399810","13398855","13397948","13397798","13396168","13394919","13393447","13393367","13392620","13392489","13392472","13392176"]

const notPay = "http://crm.soqi.cn/publicOceanUnpay/unPayOceanLock.xhtml"
const newSrc = "http://crm.soqi.cn/publicOcean/mpOceanLock.xhtml"
// 
// 
function postUid(uid) {
    options.headers['Content-Length'] = 12 + uid.length
    const req = http.request(newSrc, options, (res) => {
        console.log(`状态码: ${res.statusCode}`);
        console.log(`响应头: ${JSON.stringify(res.headers)}`);
        res.setEncoding('utf8');
        res.on('data', (chunk) => {
            console.log(`响应主体: ${chunk}`);
        });
        res.on('end', () => {
            console.log('响应中已无数据');
        });
    });

    req.on('error', (e) => {
        console.error(`请求遇到问题: ${e.message}`);
    });
    //mpId
    // recordId
    req.write(`mpId=${uid}&type=1`);
    req.end();
}

JSON.parse(data).forEach(id => {
    postUid(id)
})