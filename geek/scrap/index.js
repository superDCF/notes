const puppeteer = require('puppeteer');
const fs = require("fs")

let data_list = []
puppeteer.launch({
    headless: false,
    devtools: true,
    defaultViewport: {
        width: 1300,
        height: 800
    }
}).then(async browser => {
    const page = await browser.newPage();
    await page.goto('https://account.geekbang.org/login');
    //   await browser.close();
    // let account = await getNamePW()
    // console.log(account.name)
    await sleep(2)
    inputNamePW(page, await getNamePW(), browser)
    await sleep(2)
    // var cur_cookie = await page.cookies()
    // console.log(cur_cookie)
    await sleep(5)
    await page.click("a.ment-link:nth-child(6)")
    await sleep(8)
    // data_list
    page.on('requestfinished', async require => {
        try {
            console.log("url", await require.response().json())
        } catch (error) {
            console.log("err", error)
        }
    })
    await page.click(".btn-wrapper")
});

async function inputNamePW(page, account, browser) {
    // let input = page.$('.nw-input')
    // console.log(input.value)
    page.type('.nw-input', account.name.toString())
    await sleep()
    page.type('.input', account.pw.toString())
    await sleep()
    page.click('.mybtn')
    // await sleep()
    // let cookies = await page.cookies()
    await sleep(4)
    // page.setCookie(...cookie)
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