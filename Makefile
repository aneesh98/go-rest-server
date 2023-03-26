MONGODB_SCRIPT = ./db_population_helper/mongo/main.js

mongodb_fake_populate:
	node ${MONGODB_SCRIPT} ${SIZE}