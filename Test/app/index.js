import express from 'express';
const app = express();

const port = 3000;

// Function to log current time
const logTime = () => {
    const currentTime = new Date().toLocaleTimeString();
    console.log(`Current time: ${currentTime}`);
};

// Set up interval to log time every 5 seconds
setInterval(logTime, 1000);

app.get('/', (req, res) => {
    res.send('Dummy app running. Check console for time logs.');
});

app.listen(port, () => {
    console.log(`Dummy app listening at http://localhost:${port}`);
    console.log('Time will be logged every 5 seconds. Check console for updates.');
});
