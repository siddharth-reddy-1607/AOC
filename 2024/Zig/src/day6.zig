const std = @import("std");

fn resetSeen(seen: [][]u32) void {
    const rows = seen.len;
    const cols = seen[0].len;
    for (0..rows) |r| {
        for (0..cols) |c| {
            seen[r][c] = 0;
        }
    }
}

fn resetSeenWithDirection(seen: [][][4]u32) void {
    const rows = seen.len;
    const cols = seen[0].len;
    for (0..rows) |r| {
        for (0..cols) |c| {
            for (0..4) |dir| {
                seen[r][c][dir] = 0;
            }
        }
    }
}

fn getMovesBeforeExitOrLoop(guard: [2]i32, grid: std.ArrayList([]u8), seen: [][]u32, seenWithDirection: [][][4]u32) struct { uniqueCellsVisited: u32, loop: bool } {
    const rows = grid.items.len;
    const cols = grid.items[0].len;
    resetSeen(seen);
    resetSeenWithDirection(seenWithDirection);
    const directions = [_][2]i32{
        [2]i32{ -1, 0 },
        [2]i32{ 0, 1 },
        [2]i32{ 1, 0 },
        [2]i32{ 0, -1 },
    };
    var dirIdx: u32 = 0;
    var r: i32 = guard[0];
    var c: i32 = guard[1];
    var uniqueCellsVisited: u32 = 0;
    while (true) {
        const newR: i32 = r + directions[dirIdx][0];
        const newC: i32 = c + directions[dirIdx][1];
        if (newR < 0 or newR >= rows or newC < 0 or newC >= cols) {
            return .{ .uniqueCellsVisited = uniqueCellsVisited, .loop = false };
        }
        if (grid.items[@intCast(newR)][@intCast(newC)] == '#') {
            dirIdx = (dirIdx + 1) % 4;
            continue;
        }
        r = newR;
        c = newC;
        if (seenWithDirection[@intCast(r)][@intCast(c)][dirIdx] == 1) {
            return .{ .uniqueCellsVisited = 0, .loop = true };
        }
        seenWithDirection[@intCast(r)][@intCast(c)][dirIdx] = 1;
        if (seen[@intCast(newR)][@intCast(newC)] == 1) {
            continue;
        }
        seen[@intCast(newR)][@intCast(newC)] = 1;
        uniqueCellsVisited += 1;
    }
    return .{ .uniqueCellsVisited = 0, .loop = false };
}

fn day6_sol(allocator: std.mem.Allocator) !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();
    var grid = std.ArrayList([]u8).init(allocator);
    defer {
        for (grid.items) |row| {
            allocator.free(row);
        }
        grid.deinit();
    }
    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(usize))) |line| {
        try grid.append(line);
    }
    const rows: usize = grid.items.len;
    const cols: usize = grid.items[0].len;
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
    var seenWithDirection = try allocator.alloc([][4]u32, rows);
    for (0..rows) |r| {
        seenWithDirection[r] = try allocator.alloc([4]u32, cols);
    }
    defer {
        for (0..rows) |r| {
            allocator.free(seenWithDirection[r]);
        }
        allocator.free(seenWithDirection);
    }
    var guard = [2]i32{ 0, 0 };
    for (0..rows) |r| {
        for (0..cols) |c| {
            if (grid.items[r][c] == '^') {
                guard[0] = @intCast(r);
                guard[1] = @intCast(c);
                break;
            }
        }
    }
    const part1: u32 = getMovesBeforeExitOrLoop(guard, grid, seen, seenWithDirection).uniqueCellsVisited;
    var part2: u32 = 0;
    for (0..rows) |r| {
        for (0..cols) |c| {
            if (grid.items[r][c] == '^' or grid.items[r][c] == '#') {
                continue;
            }
            grid.items[r][c] = '#';
            if (getMovesBeforeExitOrLoop(guard, grid, seen, seenWithDirection).loop) {
                part2 += 1;
            }
            grid.items[r][c] = '.';
        }
    }
    std.debug.print("Part 1: {d}, Part 2 : {d}\n", .{ part1, part2 });
}

test "day6" {
    try day6_sol(std.testing.allocator);
}
