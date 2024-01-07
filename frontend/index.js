const express = require('express');
const path = require('path');
const app = express();

// SWITCH TO SSR
app.get('/', function (req, res) {
    res.sendFile(path.resolve(__dirname, "views/index.html"));
})

app.listen(8000);