const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

const app = express();

const pong = "PONG from js";

app.use(bodyParser.json());

app.use(cors());

app.get('/ping', (req, res) => {
  res.send(pong);
});

app.listen(8003, () => {
  console.log('listening on port 8003');
});
