import express from "express";

const app = express();
const PORT = 8081;

app.get("/temperature", (req, res) => {
    let { location, sensorId } = req.query;

    // Если location не передан — определить по sensorId
    if (!location) {
        switch (sensorId) {
            case "1":
                location = "Living Room";
                break;
            case "2":
                location = "Bedroom";
                break;
            case "3":
                location = "Kitchen";
                break;
            default:
                location = "Unknown";
        }
    }

    // Если sensorId не передан — определить по location
    if (!sensorId) {
        switch (location) {
            case "Living Room":
                sensorId = "1";
                break;
            case "Bedroom":
                sensorId = "2";
                break;
            case "Kitchen":
                sensorId = "3";
                break;
            default:
                sensorId = "0";
        }
    }

    const temperature = (Math.random() * 10 + 18).toFixed(2);

    res.json({
        sensorId,
        location,
        temperature: Number(temperature),
        unit: "°C",
        timestamp: new Date().toISOString()
    });
});

app.listen(PORT, () => {
    console.log(`Temperature API running on port ${PORT}`);
});