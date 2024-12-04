const std = @import("std");

const directions = [_][2]i32{
    [_]i32{ 1, 0 },
    [_]i32{ -1, 0 },
    [_]i32{ 0, 1 },
    [_]i32{ 0, -1 },
    [_]i32{ 1, 1 },
    [_]i32{ 1, -1 },
    [_]i32{ -1, 1 },
    [_]i32{ -1, -1 },
};

fn countXMAS(grid: *std.ArrayList(std.ArrayList(u8)), r: usize, c: usize) u32 {
    var xmas: u32 = 0;
    const rows = grid.items.len;
    const cols = grid.items[0].items.len;
    for (directions) |direction| {
        var curRow: i32 = @intCast(r);
        var curCol: i32 = @intCast(c);
        const str = "MAS";
        var found = true;
        for (1..4) |idx| {
            curRow = @addWithOverflow(curRow, direction[0])[0];
            curCol = @addWithOverflow(curCol, direction[1])[0];
            if (curRow == rows or curRow < 0 or curCol == cols or curCol < 0) {
                found = false;
                break;
            }
            if (str[idx - 1] != grid.items[@intCast(curRow)].items[@intCast(curCol)]) {
                found = false;
                break;
            }
        }
        if (found) {
            xmas += 1;
        }
    }
    return xmas;
}

fn countX_MAS(grid: *std.ArrayList(std.ArrayList(u8)), r: usize, c: usize) u32 {
    const rows = grid.items.len;
    const cols = grid.items[0].items.len;
    if (r + 1 == rows or c + 1 == cols or r == 0 or c == 0) {
        return 0;
    }
    if (grid.items[r - 1].items[c - 1] == 'M' and grid.items[r + 1].items[c + 1] == 'S' and grid.items[r + 1].items[c - 1] == 'M' and grid.items[r - 1].items[c + 1] == 'S') {
        return 1;
    }
    if (grid.items[r - 1].items[c - 1] == 'M' and grid.items[r + 1].items[c + 1] == 'S' and grid.items[r + 1].items[c - 1] == 'S' and grid.items[r - 1].items[c + 1] == 'M') {
        return 1;
    }
    if (grid.items[r - 1].items[c - 1] == 'S' and grid.items[r + 1].items[c + 1] == 'M' and grid.items[r + 1].items[c - 1] == 'S' and grid.items[r - 1].items[c + 1] == 'M') {
        return 1;
    }
    if (grid.items[r - 1].items[c - 1] == 'S' and grid.items[r + 1].items[c + 1] == 'M' and grid.items[r + 1].items[c - 1] == 'M' and grid.items[r - 1].items[c + 1] == 'S') {
        return 1;
    }
    return 0;
}
pub fn day4_sol(allocator: std.mem.Allocator) !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();
    var grid = std.ArrayList(std.ArrayList(u8)).init(allocator);
    defer grid.deinit();
    var part1: u32 = 0;
    var part2: u32 = 0;
    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(u32))) |line| {
        var row = std.ArrayList(u8).init(allocator);
        defer allocator.free(line);
        for (line) |char| {
            try row.append(char);
        }
        try grid.append(row);
    }
    const rows: usize = grid.items.len;
    const cols: usize = grid.items[0].items.len;
    for (0..rows) |r| {
        for (0..cols) |c| {
            if (grid.items[r].items[c] == 'X') {
                part1 += countXMAS(&grid, r, c);
            }
            if (grid.items[r].items[c] == 'A') {
                part2 += countX_MAS(&grid, r, c);
            }
        }
    }
    for (grid.items) |row| {
        row.deinit();
    }
    std.debug.print("Part 1 : {d}, Part 2 : {d}", .{ part1, part2 });
}
