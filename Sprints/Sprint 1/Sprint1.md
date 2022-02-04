# Sprint 1

01/21 - 02/04

## Summary

In Sprint 1, our team successfully initialized the front and backend of our application, learning a lot about TypeScript and GoLang along the way.

In the backend, we initialized the Go server and created our first database schema for an "establishment", using GORM connected to a local sqlite database. We populated our table of restaurant establishments using data found by scraping the [UF Campus Map website](https://campusmap.ufl.edu/). Using a popular Go package "Gin", we also set up our first api endpoints to get, create, and delete establishments from this table.

In the frontend, we created the foundation of our React application using the create-react-app package to initialize the project. Using MaterialUI, we began establishing the beginnings of our app layout and theme. We created and made good progress on one of our main app pages, which is the list of restaurant establishments paired with a map showing their locations (powered by Leaflet and OpenStreetMap).

## Links

[GitHub Repo](https://github.com/Monicakodali/SEPROJECT)
[GitHub Project Link](https://github.com/Monicakodali/SEPROJECT/projects/1)
[GitHub PRs for Sprint 1](https://github.com/Monicakodali/SEPROJECT/pulls?q=is%3Apr+is%3Aclosed+created%3A%3E2022-01-24+created%3A%3C2022-02-04)

## Completed Stories

\#1 - As a developer, I would like to have an initialized Go backend
\#3 - As a developer, I would like to initialize the database and connect to it from the backend 
\#5 - As a developer, I want an Establishments table created and populated in the DB
\#6 - As a developer, I want 1 basic API end point set up
\#7 - As a developer, I want a basic react app as a foundation
\#9 - As a user, I want to see a list of restaurants with info and map locations