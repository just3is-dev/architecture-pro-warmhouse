import express from "express";

const app = express();
app.use(express.json());

const devices = [];
let nextId = 1;

app.get("/devices", (req, res) => {
    res.json(devices);
});

app.get("/devices/:id", (req, res) => {
    const id = Number(req.params.id);
    const device = devices.find(d => d.id === id);

    if (!device) {
        return res.status(404).json({ error: "Device not found" });
    }

    res.json(device);
});

app.post("/devices", (req, res) => {
    const device = {
        id: nextId++,
        name: req.body.name || "Unnamed device",
        status: "offline"
    };

    devices.push(device);

    res.status(201).json(device);
});

app.put("/devices/:id", (req, res) => {
    const id = Number(req.params.id);
    const device = devices.find(d => d.id === id);

    if (!device) {
        return res.status(404).json({ error: "Device not found" });
    }

    device.name = req.body.name ?? device.name;
    device.status = req.body.status ?? device.status;

    res.json(device);
});

app.patch("/devices/:id/status", (req, res) => {
    const id = Number(req.params.id);
    const device = devices.find(d => d.id === id);

    if (!device) {
        return res.status(404).json({ error: "Device not found" });
    }

    device.status = req.body.status || device.status;

    res.json(device);
});

app.delete("/devices/:id", (req, res) => {
    const id = Number(req.params.id);
    const index = devices.findIndex(d => d.id === id);

    if (index === -1) {
        return res.status(404).json({ error: "Device not found" });
    }

    const deleted = devices.splice(index, 1);

    res.json({
        message: "Device deleted",
        device: deleted[0]
    });
});

app.listen(8082, () => {
    console.log("Device Service running on 8082");
});