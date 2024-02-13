# MPS Support in the GPU Device Plugin


An init container in the MPS control daemon daemonset creates a `tmpfs` mount at
`/run/nvidia/mps/shm`. Since `/run/nvidia/mps` is mounted with `Bidirectional` mount-propagation,
this `tmpfs` is also visible on the host.

The MPS control daemon container mounts:
* The driver root (e.g. `/`) to `/driver-root`
* The `tmpfs` at `/run/nvidia/mps/shm` to `/driver-root/dev/shm`
* The MPS root `/run/nvidia/mps` to `/driver-root/mps`

The requirement to mount these to `/driver-root` is due to the fact that we currently
run the `nvidia-cuda-mps-control` daemon after a `chroot /driver-root` to ensure that
the `nvidia-cuda-mps-server` can be found by the control daemon once it is started.

This also means that `/driver-root` *cannot be mounted read-only* as the mount
to `/driver-root/mps` then fails.

This has some side-effects:
* The `/driver-root/mps` mount creates an `mps` folder on the host -- even though this does not contain the contents created by the MPS daemon.
* The path mounted as `/driver-root` from the host is writable from the container.