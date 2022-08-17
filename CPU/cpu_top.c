#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/mm.h>
#include <linux/sched/signal.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Edgar Borrayo - 201503702");
MODULE_DESCRIPTION("Basic process information Linux module.");
MODULE_VERSION("0.1");


int proc_count(void)
{
  int i=0;
  struct task_struct *thechild;
  for_each_process(thechild){
    i++;

  }
  return i;
}

char buffer[256];
char * get_task_state(long state)
{
    switch (state) {
        case TASK_RUNNING:
            return "TASK_RUNNING";
        case TASK_INTERRUPTIBLE:
            return "TASK_INTERRUPTIBLE";
        case TASK_UNINTERRUPTIBLE:
            return "TASK_UNINTERRUPTIBLE";
        case TASK_REPORT_IDLE:
            return "TASK_REPORT_IDLE";
        case TASK_IDLE:
            return "TASK_IDLE";
        case __TASK_STOPPED:
            return "__TASK_STOPPED";
        case __TASK_TRACED:
            return "__TASK_TRACED";
        default:
        {
            sprintf(buffer, "Unknown Type:%ld\n", state);
            return buffer;
        }
    }
}

static int writeFile(struct seq_file* archivo, void *v){
    int i = 0;
    bool comma = false;
    struct task_struct *thechild;

    seq_printf(archivo, "[\n");
    for_each_process(thechild){
        
        if (comma) {
            seq_printf(
                archivo, 
                ",\n\t{\"UID\": %d, \"PID\": %d, \"Padre\": %d, \"Process\": \"%s\", \"State\": \"%s\"}", 
                thechild->real_cred->uid.val, 
                thechild->pid, 
                thechild->real_parent->pid, 
                thechild->comm, 
                get_task_state(thechild->__state)
            );
        }else{
            seq_printf(
                archivo, 
                "\t{\"UID\": %d, \"PID\": %d, \"Padre\": %d, \"Process\": \"%s\", \"State\": \"%s\"}", 
                thechild->real_cred->uid.val, 
                thechild->pid, 
                thechild->real_parent->pid, 
                thechild->comm, 
                get_task_state(thechild->__state)
            );  
            comma = true;  
        }    
        i++;
    }
    seq_printf(archivo, "\n]\n");
    pr_info("Number of processes: %u\n", i);

    return 0;
}

static int atOpen(struct inode* inode, struct file* file){
    return single_open(file, writeFile, NULL);
}

static struct proc_ops ops={
    .proc_open = atOpen,
    .proc_release = single_release,
    .proc_read = seq_read,
    .proc_lseek = seq_lseek
};

static int load_module(void) {
    printk(KERN_INFO "El grupo 2 ha montado el monitor de cpu\n");
    printk(KERN_INFO "Total running processes: %d .\n", proc_count());
    proc_create("proc.json", 0, NULL, &ops);
    return 0;
}

static void unload_module(void) {

    printk(KERN_INFO "El grupo 2 ha removido el monitor de cpu.”\n");
    remove_proc_entry("proc.json", NULL);
}

module_init(load_module);
module_exit(unload_module);