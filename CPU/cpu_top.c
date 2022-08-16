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

        seq_printf(m, "<table id=\"table\" class=\"jsTT\">");
        seq_printf(m, "<tr><th>PID</th><th>Nombre</th><th>Estado</th><th>Memoria</th><th>ACCION</th></tr>");
        
        for_each_process(tarea) {
			memoria=0;
			cpu=0;
			cpu = tarea->utime + tarea->stime;

             if(tarea->mm){ 
                 memoria = tarea->mm->total_vm * 100000 / 2048;
                // memoria = (tarea->mm->mmap->vm_end - tarea->mm->mmap->vm_start );
             }
		  
			/*if(tarea->mm){
				memoria = tarea->mm->task_size;
			}      */     

            seq_printf(m, "<tr data-tt-id=\"%d\" data-tt-level=\"0\"> <td style=\"white-space: nowrap;\">   %d  </td>",
                tarea->pid,
                tarea->pid
            );

            seq_printf(m, "<td>%s</td>",
                tarea->comm
            );

            seq_printf(m, "<td>%c</td>",
                estados(tarea->state)
            );

            seq_printf(m, "<td>%ld</td>",
                memoria
            ); 

            seq_printf(m, "<td> <button type=\"button\" onclick=\"terminarProceso(%d)\">Kill</button></td>",
            tarea->pid                
            );          

            seq_printf(m, "</tr>");

        
            list_for_each(list, &tarea->children){
            
            if(contadorHijos == 0){
                    contadorHijos = contadorHijos + 1; 
                }else{
                    contadorHijos = contadorHijos + 1;
                }
            
            task_child = list_entry( list, struct task_struct, sibling );
            
            seq_printf(m, "<tr data-tt-id=\"NA\" data-tt-parent-id=\"%d\"data-tt-level=\"1\"> <td style=\"white-space: nowrap;\"><span style=\"padding-left:0px;\"></span><span style=\"padding-left:0px;\"></span>   %d  </td>",
                tarea->pid,
                task_child->pid
            );

            seq_printf(m, "<td>%s</td>",
                task_child->comm
            );

            seq_printf(m, "<td>%c</td>",
                estados(task_child->state)
            );

            seq_printf(m, "<td>%ld</td>",
                memoria
            ); 

            seq_printf(m, "<td></td>"                
            );          

            seq_printf(m, "</tr>");

         
            }            
                if(contadorHijos > 0){
                    contadorHijos = 0;
                } 
			contador = contador + 1;
        }
            
        seq_printf(m, "</table>");    
        seq_printf(m, "<input id=\"procesoEjecucion\" name=\"procesoEjecucion\" type=\"hidden\" value=\"%d\">",
        procesoEjecucion
        );
        seq_printf(m, "<input id=\"procesoSuspendidos\" name=\"procesoSuspendidos\" type=\"hidden\" value=\"%d\">",
        procesoSuspendidos
        );
        seq_printf(m, "<input id=\"procesoDetenido\" name=\"procesoDetenido\" type=\"hidden\" value=\"%d\">",
        procesoDetenido
        );
        seq_printf(m, "<input id=\"procesoZombie\" name=\"procesoZombie\" type=\"hidden\" value=\"%d\">",
        procesoZombie
        );
        seq_printf(m, "<input id=\"procesosTotales\" name=\"procesosTotales\" type=\"hidden\" value=\"%d\">",
        procesosTotales
        );
        return 0;
    }
       
    static int meminfo_proc_open(struct inode *inode, struct file *file)
    {
        return single_open(file, meminfo_proc_show, NULL);
    }

    static const struct file_operations meminfo_proc_fops = {
        .open       = meminfo_proc_open,
        .read       = seq_read,
        .llseek     = seq_lseek,
        .release    = single_release,
    };

    
    static int __init inicio_mod(void)
    {
		//Este metodo escribe el archivo en la carpeta PROC
        printk(KERN_INFO "Nombre: Grupo2 \n");
        proc_create("cpu_top", 0, NULL, &meminfo_proc_fops);
        return 0;
    }

	static void __exit final_mod(void)
    {   
		//Este metodo sale del modulo
        printk(KERN_INFO "Descargando: ModuloG2 \n");
        remove_proc_entry("cpu_top",NULL);
    }

    module_init(inicio_mod);
    module_exit(final_mod);

	MODULE_LICENSE("GPL");