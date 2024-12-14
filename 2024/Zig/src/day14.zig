const std = @import("std");

const maxX: i32 = 101;
const maxY: i32 = 103;

fn getNewPosition(position: [2]i32, velocity: [2]i32, t: i32) [2]i32 {
    return [2]i32{ @mod((position[0] + velocity[0] * t + maxX * t), maxX), @mod((position[1] + velocity[1] * t + maxY * t), maxY) };
}

fn getSafetyScore(positions: std.ArrayList([2]i32), velocities: std.ArrayList([2]i32), t: i32, newPositions: *std.ArrayList([2]i32)) !u32 {
    newPositions.clearRetainingCapacity();
    const length = positions.items.len;
    var q1: u32 = 0;
    var q2: u32 = 0;
    var q3: u32 = 0;
    var q4: u32 = 0;
    for (0..length) |idx| {
        const newPos = getNewPosition(positions.items[idx], velocities.items[idx], t);
        try newPositions.append(newPos);
        if (newPos[0] < maxX / 2 and newPos[1] < maxY / 2) {
            q1 += 1;
        }
        if (newPos[0] > maxX / 2 and newPos[1] < maxY / 2) {
            q2 += 1;
        }
        if (newPos[0] < maxX / 2 and newPos[1] > maxY / 2) {
            q3 += 1;
        }
        if (newPos[0] > maxX / 2 and newPos[1] > maxY / 2) {
            q4 += 1;
        }
    }
    return q1 * q2 * q3 * q4;
}

fn printGrid(allocator: std.mem.Allocator, positions: std.ArrayList([2]i32)) !void {
    var grid = try allocator.alloc([]u8, maxY + 1);
    for (0..maxY) |y| {
        grid[y] = try allocator.alloc(u8, maxX + 1);
    }
    defer {
        for (0..maxY) |y| {
            allocator.free(grid[y]);
        }
        allocator.free(grid);
    }
    for (0..maxY) |y| {
        for (0..maxX) |x| {
            grid[y][x] = ' ';
        }
    }
    for (positions.items) |pos| {
        grid[@intCast(pos[1])][@intCast(pos[0])] = '*';
    }
    for (0..maxY) |y| {
        std.debug.print("{s}\n", .{grid[y]});
    }
}

fn day14_sol(allocator: std.mem.Allocator) !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();
    var positions = std.ArrayList([2]i32).init(allocator);
    defer positions.deinit();
    var velocities = std.ArrayList([2]i32).init(allocator);
    defer velocities.deinit();
    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(u64))) |line| {
        defer allocator.free(line);
        var spaceIt = std.mem.splitScalar(u8, line, ' ');
        var idx: u32 = 0;
        while (spaceIt.next()) |token| {
            if (idx == 0) {
                var pos = std.mem.splitScalar(u8, std.mem.trimLeft(u8, token, "p="), ',');
                try positions.append([2]i32{ try std.fmt.parseInt(i32, pos.next().?, 10), try std.fmt.parseInt(i32, pos.next().?, 10) });
                idx += 1;
            } else {
                var vel = std.mem.splitScalar(u8, std.mem.trim(u8, token, "v="), ',');
                try velocities.append([2]i32{ try std.fmt.parseInt(i32, vel.next().?, 10), try std.fmt.parseInt(i32, vel.next().?, 10) });
            }
        }
    }
    var newPostions = std.ArrayList([2]i32).init(allocator);
    defer newPostions.deinit();
    const part1: u32 = try getSafetyScore(positions, velocities, 100, &newPostions);
    std.debug.print("Part 1: {d}\n", .{part1});
    //Already implemented in Go, so I know the range here. The safety score is low for the case where we have an XMAS tree as the points be close together. Total 500 points
    //Trail and error on the safety score
    for (7500..8000) |i| {
        const safetyScore = try getSafetyScore(positions, velocities, @intCast(i), &newPostions);
        if (safetyScore <= 110 * 110 * 100 * 100) {
            std.debug.print("--------------------------------------------------Time = {d}--------------------------------------------------\n", .{i});
            try printGrid(allocator, newPostions);
        }
    }
}

test "day14" {
    try day14_sol(std.testing.allocator);
}
