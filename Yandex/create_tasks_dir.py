import os
import click

@click.command()
@click.argument('task_count', type=int)
def main(task_count):
    print('Start create directorys for ' + str(task_count) + ' tasks')
    dir_name = './tasks'
    if os.path.exists(dir_name):
        print(f'Directory {dir_name} already exists')
    else:
        os.mkdir(dir_name)
    
    for i in range(1, task_count+1):
        task_dir_name = '/task' + str(i)
        task_name = '/task' + str(i)
        if os.path.exists(f'{dir_name}{task_dir_name}'):
            print(f'Direcory {dir_name}{task_dir_name} already exists')
        else: 
            os.mkdir(f'{dir_name}{task_dir_name}')

        if os.path.exists(f'{dir_name}{task_dir_name}{task_name}.go'):
            print(f'File {dir_name}{task_dir_name}{task_name}.go already exists')
        else:
            open(f'{dir_name}{task_dir_name}{task_name}.go', 'w').close()

    print('Finish create directorys for ' + str(task_count) + ' tasks')

if __name__ == '__main__':
    main()
