import os
import click 

@click.command()
@click.argument('task_number', type=int)
def main(task_number):
    print('Start create directorys for task ' + str(task_number))
    task_name = './task' + str(task_number)
    os.mkdir(task_name)
    open(f'{task_name}/{task_name}.go', 'w').close()
    print('Finish create directorys for task ' + str(task_number))

if __name__ == '__main__':
    main()