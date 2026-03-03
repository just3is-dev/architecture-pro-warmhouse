import express from "express";
import { v4 as uuid } from "uuid";

const app = express();
app.use(express.json());

const devices = [];

app.get("/devices", (req, res) => {
    res.json(devices);
});

app.post("/devices", (req, res) => {
    const device = {
        id: uuid(),
        name: req.body.name,
        status: "offline"
    };
    devices.push(device);
    res.status(201).json(device);
});

app.patch("/devices/:id/status", (req, res) => {
    const device = devices.find(d => d.id === req.params.id);
    if (!device) return res.status(404).json({ error: "Not found" });

    device.status = req.body.status;
    res.json(device);
});

app.listen(8082, () => {
    console.log("Device Service running on 8082");
});