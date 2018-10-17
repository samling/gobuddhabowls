# First-Time Setup/New Migrations

1. `docker-compose --file docker-compose.db.yml build`
2. `docker-compose up`
3. In `../buddhabowls-data/seeddb`: `python migrate_to_postgres.py`

# Running the Application

1. `docker-compose build`
2. `docker-compose up`

The application will be available at `http://localhost:3000`
