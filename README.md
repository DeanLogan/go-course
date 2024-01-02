# Go-Course Repository

Welcome to the Go-Course repository! This repository contains code developed while following the Go programming language course on FreeCodeCamp's YouTube channel. You can find the course [here](https://www.youtube.com/watch?v=un6ZyFkqFKo&amp;list=WL&amp;index=49&amp;t=161s).

## Repository Structure

The repository is meticulously organized into two primary folders:

1. **course**: This section encompasses the code corresponding to each chapter within the instructional segment of the course.

2. **project**: Here, you'll discover the code associated with the hands-on project developed during the course.

Feel free to explore these folders and delve into the code, enhancing your understanding of Go programming.

# Course Description

The course is divided into different chapters within the [YT video](https://www.youtube.com/watch?v=un6ZyFkqFKo&amp;list=WL&amp;index=49&amp;t=161s). The concepts are explained within this video by the instructor, the code within the files for each chapter are a combination of the code created by the intstructor and some of my own notes. 

# Project: RSS Feed Aggregator

The project is a RSS feed aggregator. It's a web server that allows clients to:

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

## Getting Started

To get started with this project, clone the repository and navigate to the project directory.

### Prerequisites

- Go 1.21.5 or later
- PostgreSQL database

### Environment Variables

The project requires the following environment variables:

- `PORT`: The port on which the server will run
- `DB_URL`: The connection string for the PostgreSQL database

You can set these in a `.env` file at the root of the project directory.

### Project Structure

The project is organized into several Go files, each responsible for a specific part of the functionality:

- `main.go`: The entry point of the application. It sets up the server and the routes.
- `handler_user.go`: Contains handlers for user-related routes.
- `handler_feed.go`: Contains handlers for feed-related routes.
- `handler_feed_follows.go`: Contains handlers for feed follow-related routes.
- `handler_readiness.go`: Contains the readiness probe handler.
- `models.go`: Defines the data models used in the project.
- `rss.go`: Contains the logic for fetching and parsing RSS feeds.
- `middleware_auth.go`: Contains the authentication middleware.

Additionally, the project includes a `/sql` folder:

- `/sql`: Contains SQL scripts for database setup and migrations. These scripts are used to create the necessary tables and relationships in the database.
