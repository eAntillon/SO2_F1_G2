obj-m += cpu_top.o
all:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules
clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean

test:
	sudo dmesg -C
	sudo insmod cpu_top.ko
	sudo rmmod cpu_top.ko
	dmesg
