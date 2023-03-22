const express = require('express')
const app = express()
const ejs = require('ejs')
const path = require('path')
const history = require('connect-history-api-fallback');
const port = 3000
app.use(history())
app.use(express.static(path.join(__dirname, 'tides-ui')));
app.set("views", __dirname + "/tides-ui");
app.engine('html',ejs.__express)
// 配置模版引擎
app.set("view engine", "html")
app.get('/', (req, res) => {
  res.render('index', {})
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
