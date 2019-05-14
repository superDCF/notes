const puppeteer = require('puppeteer');
const fs = require("fs")

const cookie = [{
    name: 'SERVERID',
    value: '1fa1f330efedec1559b3abbcb6e30f50|1557807172|1557807170',
    domain: 'account.geekbang.org',
    path: '/',
    expires: -1,
    size: 62,
    httpOnly: false,
    secure: false,
    session: true
},
{
    name: 'GCID',
    value: '532984f-4b3bddb-a33cb22-15c4b43',
    domain: '.geekbang.org',
    path: '/',
    expires: 1560399170.920662,
    size: 35,
    httpOnly: true,
    secure: false,
    session: false
},
{
    name: 'GCESS',
    value:
        'BAIEREDaXAgBAwMEREDaXAUEAAAAAAcE_iKnjgQEAC8NAAYEGmao3QsCBAAMAQEBBDC6FQAJAQEKBAAAAAA-',
    domain: '.geekbang.org',
    path: '/',
    expires: 1558671173.077807,
    size: 89,
    httpOnly: true,
    secure: false,
    session: false
},
{
    name: '_gat',
    value: '1',
    domain: '.geekbang.org',
    path: '/',
    expires: 1557807231,
    size: 5,
    httpOnly: false,
    secure: false,
    session: false
},
{
    name: '_ga',
    value: 'GA1.2.1808592929.1557807171',
    domain: '.geekbang.org',
    path: '/',
    expires: 1620879171,
    size: 30,
    httpOnly: false,
    secure: false,
    session: false
},
{
    name: '_gid',
    value: 'GA1.2.1059326035.1557807171',
    domain: '.geekbang.org',
    path: '/',
    expires: 1557893571,
    size: 31,
    httpOnly: false,
    secure: false,
    session: false
}]

puppeteer.launch({
    headless: false,
    defaultViewport: {
        width: 1300,
        height: 800
    }
}).then(async browser => {
    const page = await browser.newPage();
    page.setCookie(...cookie)
    await page.goto('https://account.geekbang.org/login');
    //   await browser.close();
    // let account = await getNamePW()
    // console.log(account.name)
    inputNamePW(page, await getNamePW(), browser)
});

async function inputNamePW(page, account, browser) {
    // let input = page.$('.nw-input')
    // console.log(input.value)
    page.type('.nw-input', account.name.toString())
    await sleep()
    page.type('.input', account.pw.toString())
    await sleep()
    page.click('.mybtn')
    await sleep()
    let cookies = await page.cookies()
    console.log(cookies)
    await sleep(4)
    page.setCookie(...cookie)
    // const buyPage = browser.newPage()
    await sleep(4)
    // await page.goto("https://account.geekbang.org/dashboard/buy")
    // page.click('.ment-link')
}

async function getNamePW() {
    let account = await fs.readFileSync(__dirname + '/../account.json', {
        encoding: 'utf8'
    })
    // console.log(account, typeof account, typeof JSON.parse(account), )
    return JSON.parse(account)
}

function sleep(duration = 0.5) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve()
        }, duration * 1000)
    })
}