//! By convention, main.zig is where your main function lives in the case that
//! you are building an executable. If you are making a library, the convention
//! is to delete this file and start with root.zig instead.
const day = @import("day2.zig");
pub fn main() !void {
    try day.day2_sol();
}
