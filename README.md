# task-scheduler

This is a simple task scheduler that consisted of two programs:

1. The main program of a scheduler bundled in an API with features such as:
   * Starting the scheduler
   * Stopping the scheduler
   * Adding task to the scheduler
   * Deleting existing task from the scheduler
   * Executing the task and function at each specified time interval

    This program currently only supports two example executable functions (noted by the option number, which needs to be correctly inputted), as a demonstration that the program can execute various function:
      1. Printing the task name
      2. Printing the task name with the timestamp and the next execution timestamp

    The folder structure of this program utilize clean architecture, where there are separation of layers between each function so that every of one them has a specific purpose

    To run the program, use the command `make run-api`

2. The secondary program of a CLI app that interacts with the API so you can add or delete task via CLI

    For the option of adding task:
    - Make sure that each time interval field is inputted with the proper **crontab** format, as shown [as you can try it here](https://crontab.guru)
    - Make sure that the inputted function option is supported by the program, so the numerical input can only be within the range of 1-2.
    
    For the option of deleting task, make sure that the inputted task ID to be a positive integer and the task ID has been entered to the task scheduler
   
    To run the program, use the command `make run-cli` (Make sure that the first program have been running before starting the CLI program)
