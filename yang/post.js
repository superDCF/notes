const querystring = require("querystring");
const http = require("http");
const fs = require("fs");
const options = {
    method: "POST",
    headers: {
        "Content-Type": "application/x-www-form-urlencoded",
        Cookie:
            "OUTFOX_SEARCH_USER_ID_NCOO=365594159.8071182; Identify=AAE1399D1B6D7288B64490A9DD58479955C750A8CB3504A4B5D9335575B0B441573DFFB33C6B3463; acw_tc=76b20f6816381125179606287e5cdf3562de9ecff6dc805b072b2b4bedf1f6; radius=101.232.73.202; uudid=cms5c3bc48c-ea65-2c99-9260-1b9140b334bf; JSESSIONID=976AEC5DB82E58AB8D78D90F6D73A53C; Mcc=4F13C1F88FE88E0D7CB0BCD31B48A3E7",
        Host: "crm.soqi.cn",
        Referer: "http://crm.soqi.cn/baibao/index.html",
        "User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
        Origin: "http://crm.soqi.cn",
        Accept: "application/json, text/plain, */*",
    },
};

// 将数据写入请求主体。
let data = fs.readFileSync("./data.json", "utf8");
console.log(JSON.parse(data));
console.log(typeof JSON.parse(data));
console.log(JSON.parse(data)[0]);

// const uid = ["13438137","13436295","13436107","13435810","13435776","13435777","13427510","13427123","13427031","13426858","13426561","13426508","13426474","13426453","13426096","13425639","13425272","13424862","13424843","13424055","13423895","13423518","13422230","13422182","13422028","13422022","13421727","13421478","13420841","13420787","13420663","13420307","13420297","13420179","13419853","13419838","13419802","13419712","13419666","13419143","13418197","13417510","13416768","13415790","13415501","13415502","13415440","13415348","13415258","13415081","13414161","13412777","13412743","13412568","13410930","13410855","13410658","13410626","13410594","13410390","13410196","13409875","13408603","13408540","13408402","13408117","13408104","13408013","13407969","13407961","13407931","13407882","13407793","13407740","13407705","13406865","13406754","13406589","13406542","13406153","13405395","13404845","13404664","13404117","13403776","13402351","13401704","13400944","13399810","13398855","13397948","13397798","13396168","13394919","13393447","13393367","13392620","13392489","13392472","13392176"]

const notPay = "http://crm.soqi.cn/publicOceanUnpay/unPayOceanLock.xhtml";
const newSrc = "http://crm.soqi.cn/publicOcean/mpOceanLock.xhtml";
//
//
function postUid(uid) {
    options.headers["Content-Length"] = `mpId=${uid}&type=1`.length;
    const req = http.request(newSrc, options, (res) => {
        console.log(`状态码: ${res.statusCode}`);
        console.log(`响应头: ${JSON.stringify(res.headers)}`);
        res.setEncoding("utf8");
        res.on("data", (chunk) => {
            console.log(`响应主体: ${chunk}`);
        });
        res.on("end", () => {
            console.log("响应中已无数据");
        });
    });

    req.on("error", (e) => {
        console.error(`请求遇到问题: ${e.message}`);
    });
    //mpId
    // recordId
    req.write(`mpId=${uid}&type=1`);
    req.end();
}

JSON.parse(data).forEach((id) => {
    postUid(id);
});
