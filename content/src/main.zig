const tic = @import("tic");

export fn TIC() void {
    tic.cls(6);
    if (tic.btn(4)) tic.exit();
}
