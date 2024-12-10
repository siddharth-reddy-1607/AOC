const std = @import("std");

const directions = [4][2]i32{
    [2]i32{ 1, 0 },
    [2]i32{ -1, 0 },
    [2]i32{ 0, 1 },
    [2]i32{ 0, -1 },
};

fn getTrailHeadRating(row: usize, col: usize, grid: std.ArrayList([]u8)) u32 {
    if (grid.items[row][col] == '9') {
        return 1;
    }
    const rows: i32 = @intCast(grid.items.len);
    const cols: i32 = @intCast(grid.items[0].len);
    const r: i32 = @intCast(row);
    const c: i32 = @intCast(col);
    var ways: u32 = 0;
    for (directions) |direction| {
        const newR: i32 = r + direction[0];
        const newC: i32 = c + direction[1];
        if (newR < 0 or newR >= rows or newC < 0 or newC >= cols) {
            continue;
        }
        const diff: i32 = @subWithOverflow((grid.items[@intCast(newR)][@intCast(newC)] - '0'), (grid.items[@intCast(r)][@intCast(c)] - '0'))[0];
        if (diff != 1) {
            continue;
        }
        ways += getTrailHeadRating(@intCast(newR), @intCast(newC), grid);
    }
    return ways;
}

fn getTrailHeadScore(allocator: std.mem.Allocator, row: usize, col: usize, grid: std.ArrayList([]u8)) !u32 {
    const rows: i32 = @intCast(grid.items.len);
    const cols: i32 = @intCast(grid.items[0].len);
    var score: u32 = 0;
    var queue = std.ArrayList([2]usize).init(allocator);
    defer queue.deinit();
    var queueIdx: usize = 0;
    try queue.append([2]usize{ row, col });
    var seen = try allocator.alloc([]usize, @intCast(rows));
    for (0..@intCast(rows)) |r| {
        seen[r] = try allocator.alloc(usize, @intCast(cols));
    }
    for (0..@intCast(rows)) |r| {
        for (0..@intCast(cols)) |c| {
            seen[r][c] = 0;
        }
    }
    defer {
        for (0..@intCast(rows)) |r| {
            allocator.free(seen[r]);
        }
        allocator.free(seen);
    }
    while (queueIdx < queue.items.len) {
        const pos = queue.items[queueIdx];
        const r: i32 = @intCast(pos[0]);
        const c: i32 = @intCast(pos[1]);
        queueIdx += 1;
        if (grid.items[@intCast(r)][@intCast(c)] == '9') {
            score += 1;
            continue;
        }
        for (directions) |direction| {
            const newR: i32 = r + direction[0];
            const newC: i32 = c + direction[1];
            if (newR < 0 or newR >= rows or newC < 0 or newC >= cols) {
                continue;
            }
            const diff: i32 = @subWithOverflow((grid.items[@intCast(newR)][@intCast(newC)] - '0'), (grid.items[@intCast(r)][@intCast(c)] - '0'))[0];
            if (diff != 1) {
                continue;
            }
            if (seen[@intCast(newR)][@intCast(newC)] == 1) {
                continue;
            }
            seen[@intCast(newR)][@intCast(newC)] = 1;
            try queue.append([2]usize{ @intCast(newR), @intCast(newC) });
        }
    }
    return score;
}
fn day10_sol(allocator: std.mem.Allocator) !void {
    var part1: u32 = 0;
    var part2: u32 = 0;
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();
    var grid = std.ArrayList([]u8).init(allocator);
    defer {
        for (grid.items) |row| {
            allocator.free(row);
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
            if (grid.items[r][c] == '0') {
                part1 += try getTrailHeadScore(allocator, r, c, grid);
                part2 += getTrailHeadRating(r, c, grid);
            }
        }
    }
    std.debug.print("Part 1 : {d},Part 2 : {d}\n", .{ part1, part2 });
}

test "day10" {
    try day10_sol(std.testing.allocator);
}
