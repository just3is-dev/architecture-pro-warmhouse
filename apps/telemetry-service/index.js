import express from "express";

const app = express();
app.use(express.json());

const telemetry = [];

app.post("/telemetry", (req, res) => {
    const record = {
        deviceId: req.body.deviceId,
        value: req.body.value,
        timestamp: new Date().toISOString()
    };

    telemetry.push(record);
    res.status(201).json(record);
});

app.get("/telemetry", (req, res) => {
    const { deviceId } = req.query;
    const result = telemetry.filter(t => t.deviceId === deviceId);
    res.json(result);
});

app.listen(8083, () => {
    console.log("Telemetry Service running on 8083");
});