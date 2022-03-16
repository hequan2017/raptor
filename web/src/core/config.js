/**
 * 网站配置文件
 */

const config = {
  appName: 'raptor',
  appLogo: 'https://www.raptor.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用raptor，开源地址`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:V2.5.0`
      )
    )
    console.log(
      chalk.green(
        `> `
      )
    )
    console.log(
      chalk.green(
        `> `
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log('\n')
  }
}

export default config
