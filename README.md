# Task Manager App for CVWO Application 2022

## Task
The task for this project is to build a fullstack task manager app, and was done over the winter break.

The target audience is self defined, where in this case it would be for the general public.

Some features being implemented include:

- Tag System (Users are able to assign a tag to their tasks)
- Filter out the tasks by tags
- Update status of tasks and delete them
- User authentication 

## Libraries + Languages
**Frontend:** ReactJS

**Backend:** Go

**Database:** MySQL

Go libraries used include:

- GORM (For MySQL Connection)
- Fiber (API Library that functions similar to Express JS)
- bcrpyt (Password Hashing)
- JWT (JSON Web Tokens)

## Project Demo
The project is hosted on Netlify. You can find it [here](https://cvwo-task-manager.netlify.app/).

The backend is hosted on Heroku. To view the frontend code, click [here](https://github.com/RyanCheungJF/CVWO-Frontend).

Here is a brief screenshot of the main page. More screenshots of examples of the different pages can be found in the `imgs` folder.

![](/imgs/mainpage.png)

## Batch Files
Unforutunately, I can't run CRON locally as I do not own a Linux subsystem. Thus, I have turned to Task Scheduler and `batch` files on Windows as an alternative. 

It is used to login to the remote database and perform a backup periodically, saving to my local `mysqldump`.

## Lessons Learnt
- REST APIs on Go using Fiber, Gorilla Mux
- JWTs
- MySQL and GORM
- Hosting on Heroku with Go and MySQL
