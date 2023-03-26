/* main.js */

// require the necessary libraries
const faker = require("@faker-js/faker").faker;
const MongoClient = require("mongodb").MongoClient;

function randomIntFromInterval(min, max) { // min and max included 
    return Math.floor(Math.random() * (max - min + 1) + min);
}

async function seedDB() {
    // Connection URL
    const uri = "mongodb://localhost:27017";

    const client = new MongoClient(uri, {
        useNewUrlParser: true,
        // useUnifiedTopology: true,
    });

    try {
        await client.connect();
        console.log("Connected correctly to server");

        const collection = client.db("ToDoListDev").collection("todoLists");

        // The drop() command destroys all data from a collection.
        // Make sure you run it against proper database and collection.
        collection.drop();

        // make a bunch of time series data
        let todoListData = [];

        for (let i = 0; i < parseInt(process.argv[2]); i++) {
            const taskName = faker.word.verb() + " " + faker.word.noun()
            const id = i + 1
            let newTask = {
                id: id,
                text: taskName,
                tags: [],
                due: faker.date.past()
            };

            for (let j = 0; j < faker.datatype.number({max: 4}); j++) {
                newTask.tags.push(faker.word.adjective())
            }
            todoListData.push(newTask);
        }
        await collection.insertMany(todoListData);

        console.log("Database seeded! :)");
    } catch (err) {
        console.log(err.stack);
    } finally {
        await client.close()
    }
}

seedDB();