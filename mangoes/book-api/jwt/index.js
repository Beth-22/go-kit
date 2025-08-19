const express = require('express');
const router = require('./routes/login')

const app = express();
app.use(express.static('./public'))
app.use(express.json()); //needs to be placed before router

app.use(router);

const port= process.env.PORT || 3000;

app.listen(port);