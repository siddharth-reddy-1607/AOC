const std = @import("std");

const isInvalidAntinodeError = error{
    OutOfBounds,
    AlreadySeen,
};

fn resetSeen(seen: [][]u32) void {
    const rows = seen.len;
    const cols = seen[0].len;
    for (0..rows) |r| {
        for (0..cols) |c| {
            seen[r][c] = 0;
        }
    }
}

fn checkAntinode(r: i32, c: i32, seen: [][]u32, grid: std.ArrayList([]u8)) isInvalidAntinodeError!u32 {
    const rows = grid.items.len;
    const cols = grid.items[0].len;
    if (r < 0 or r >= rows or c < 0 or c >= cols) {
        return isInvalidAntinodeError.OutOfBounds;
    }
    if (seen[@intCast(r)][@intCast(c)] == 1) {
        return isInvalidAntinodeError.AlreadySeen;
    }
    seen[@intCast(r)][@intCast(c)] = 1;
    return 1;
}

fn getValidAntinodes(a: [2]usize, b: [2]usize, seen: [][]u32, grid: std.ArrayList([]u8), withoutDistance: bool) u32 {
    const r1: i32 = @intCast(a[0]);
    const c1: i32 = @intCast(a[1]);
    const r2: i32 = @intCast(b[0]);
    const c2: i32 = @intCast(b[1]);
    const rowDist: i32 = @intCast(@abs(r1 - r2));
    const colDist: i32 = @intCast(@abs(c1 - c2));
    var validAntinodes: u32 = 0;
    var r: i32 = 0;
    var c: i32 = 0;
    var i: i32 = undefined;

    if (a[0] < b[0]) {
        if (a[1] < b[1]) {
            i = 1;
            while (true) : (i += 1) {
                r = r1 - i * rowDist;
                c = c1 - i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
            i = 1;
            while (true) : (i += 1) {
                r = r2 + i * rowDist;
                c = c2 + i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
        } else {
            i = 1;
            while (true) : (i += 1) {
                r = r1 - i * rowDist;
                c = c1 + i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
            i = 1;
            while (true) : (i += 1) {
                r = r2 + i * rowDist;
                c = c2 - i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
        }
    } else {
        if (a[1] < b[1]) {
            i = 1;
            while (true) : (i += 1) {
                r = r1 + i * rowDist;
                c = c1 - i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
            i = 1;
            while (true) : (i += 1) {
                r = r2 - i * rowDist;
                c = c2 + i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
        } else {
            i = 1;
            while (true) : (i += 1) {
                r = r1 + i * rowDist;
                c = c1 + i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
            i = 1;
            while (true) : (i += 1) {
                r = r2 - i * rowDist;
                c = c2 - i * colDist;
                validAntinodes += checkAntinode(r, c, seen, grid) catch |err| switch (err) {
                    isInvalidAntinodeError.OutOfBounds => break,
                    else => 0,
                };
                if (!withoutDistance) {
                    break;
                }
            }
        }
    }
    return validAntinodes;
}

fn day8_sol(allocator: std.mem.Allocator) !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();
    var grid = std.ArrayList([]u8).init(allocator);
    var hashmap = std.AutoHashMap(u8, std.ArrayList([2]usize)).init(allocator);
    defer {
        var valIt = hashmap.valueIterator();
        while (valIt.next()) |list| {
            list.deinit();
        }
        hashmap.deinit();
    }
    defer {
        for (grid.items) |item| {
            allocator.free(item);
        }
        grid.deinit();
    }
    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(u64))) |line| {
        try grid.append(line);
    }
    const rows: usize = grid.items.len;
    const cols: usize = grid.items[0].len;
    for (0..rows) |r| {
        for (0..cols) |c| {
            if (grid.items[r][c] == '.') {
                continue;
            }
            const gop = try hashmap.getOrPut(grid.items[r][c]);
            if (gop.found_existing) {
                try gop.value_ptr.*.append(.{ r, c });
            } else {
                gop.value_ptr.* = std.ArrayList([2]usize).init(allocator);
                try gop.value_ptr.*.append(.{ r, c });
            }
        }
    }
    var seen = try allocator.alloc([]u32, rows);
    for (0..rows) |r| {
        seen[r] = try allocator.alloc(u32, cols);
    }
    defer {
        for (0..rows) |r| {
            allocator.free(seen[r]);
        }
        allocator.free(seen);
    }
    var part1: u32 = 0;
    var part2: u32 = 0;
    var it = hashmap.iterator();
    while (it.next()) |pair| {
        const length = pair.value_ptr.*.items.len;
        for (0..length) |i| {
            for (i + 1..length) |j| {
                part1 += getValidAntinodes(pair.value_ptr.*.items[i], pair.value_ptr.*.items[j], seen, grid, false);
            }
        }
    }
    resetSeen(seen);
    std.debug.print("With true\n", .{});
    it = hashmap.iterator();
    while (it.next()) |pair| {
        const length = pair.value_ptr.*.items.len;
        for (0..length) |i| {
            for (i + 1..length) |j| {
                part2 += getValidAntinodes(pair.value_ptr.*.items[i], pair.value_ptr.*.items[j], seen, grid, true);
            }
        }
    }
    for (0..rows) |r| {
        for (0..cols) |c| {
            if (grid.items[r][c] != '.' and seen[r][c] != 1) {
                part2 += 1;
            }
        }
    }

    std.debug.print("Part 1: {d}, Part 2 : {d}\n", .{ part1, part2 });
}

test "day8" {
    try day8_sol(std.testing.allocator);
}
