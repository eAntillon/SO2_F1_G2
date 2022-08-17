#include <linux/fs.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/sched.h>
#include <linux/sched/signal.h>
#include <linux/module.h>
#include <linux/seq_file.h>
#include <linux/proc_fs.h>

int procesoEjecucion = 0;
int procesoSuspendidos = 0;
int procesoDetenido = 0;
int procesoZombie = 0;
int procesosTotales = 0;


char estados(volatile long estado )
{   
    procesosTotales = procesosTotales + 1;
    
    if(estado == 0)
    {
        procesoEjecucion = procesoEjecucion + 1;
        return 'R';
    }

    if(estado== 1)
    {
        procesoSuspendidos = procesoSuspendidos + 1;
        return 'S';
    }
    
    if(estado== 2)
    {
        procesoSuspendidos = procesoSuspendidos + 1;
        return 'D';
    }

    if(estado== 4)
    {
        procesoZombie = procesoZombie + 1;
        return 'Z';
    }

    if(estado== 8)
    {
        procesoDetenido = procesoDetenido + 1;
        return 'T';
    }

    if(estado == 32)
    {            
        return 'X';
    }

    procesoEjecucion = procesoEjecucion + 1;
    return 'E';
}

static int meminfo_proc_show(struct seq_file *m, void *v){
    int contador = 0;
    long memoria = 0;
    long cpu = 0;
    int contadorHijos = 0; //0: no coloque tag, 1: ya coloque tag

    struct task_struct *tarea;          //El task struct es para los procesos, deberia haber algo similar para memoria y cpu
    struct task_struct *task_child;     /*    Structure needed to iterate through task children    */
    struct list_head *list;             /*    Structure needed to iterate through the list in each task->children struct    */

    for_each_process(tarea) {
        memoria=0;
        cpu=0;
        cpu = tarea->utime + tarea->stime;

            if(tarea->mm){ 
                memoria = tarea->mm->total_vm * 100000 / 2048;
            }

        list_for_each(list, &tarea->children){
        
        if(contadorHijos == 0){
                contadorHijos = contadorHijos + 1; 
            }else{
                contadorHijos = contadorHijos + 1;
            }
        
        task_child = list_entry( list, struct task_struct, sibling );

        estados(task_child->__state);
        
        }            
            if(contadorHijos > 0){
                contadorHijos = 0;
            } 
        contador = contador + 1;
    }

    seq_printf(m, "\n\tProcesos en ejecucion: %d", procesoEjecucion);
    seq_printf(m, "\n\tProcesos suspendidos %d", procesoSuspendidos);
    seq_printf(m, "\n\tProcesos detenidos: %d", procesoDetenido);
    seq_printf(m, "\n\tProcesos Zombies: %d", procesoZombie);
    seq_printf(m, "\n\tTotal de procesos: %d", procesosTotales);

    
    return 0;
}
    
static int meminfo_proc_open(struct inode *inode, struct file *file)
{
    return single_open(file, meminfo_proc_show, NULL);
}

static const struct proc_ops meminfo_proc_ops = {
    .proc_open       = meminfo_proc_open,
    .proc_read       = seq_read,
    .proc_lseek     = seq_lseek,
    .proc_release    = single_release,
};


static int __init inicio_mod(void)
{
    //Este metodo escribe el archivo en la carpeta PROC
    printk(KERN_INFO "El grupo 2 ha montado el monitor de cpu\n");
    printk(KERN_INFO "Nombre: Practica1_G2 \n");
    proc_create("cpu_top", 0, NULL, &meminfo_proc_ops);
    return 0;
}

static void __exit final_mod(void)
{   
    //Este metodo sale del modulo
    printk(KERN_INFO "Descargando: Practica1_G2 \n");
    printk(KERN_INFO "El grupo 2 ha removido el monitor de cpu. \n");
    remove_proc_entry("cpu_top",NULL);
}

module_init(inicio_mod);
module_exit(final_mod);

MODULE_LICENSE("GPL");