const std = @import("std");

fn concat(allocator: std.mem.Allocator, a: u64, b: u64) !u64 {
    const aStr = try std.fmt.allocPrint(allocator, "{d}", .{a});
    const bStr = try std.fmt.allocPrint(allocator, "{d}", .{b});
    const concatenatedStr = try allocator.alloc(u8, aStr.len + bStr.len);
    std.mem.copyForwards(u8, concatenatedStr, aStr);
    std.mem.copyForwards(u8, concatenatedStr[aStr.len..], bStr);
    defer allocator.free(concatenatedStr);
    defer allocator.free(aStr);
    defer allocator.free(bStr);
    return try std.fmt.parseInt(u64, concatenatedStr, 10);
}

fn isPossibleWithConcat(allocator: std.mem.Allocator, idx: usize, a: u64, b: u64, arrayList: *std.ArrayList(u64), target: u64) !bool {
    if (idx == arrayList.items.len) {
        return a + b == target or a * b == target or try concat(allocator, a, b) == target;
    }
    return try isPossibleWithConcat(allocator, idx + 1, a + b, arrayList.items[idx], arrayList, target) or
        try isPossibleWithConcat(allocator, idx + 1, a * b, arrayList.items[idx], arrayList, target) or
        try isPossibleWithConcat(allocator, idx + 1, try concat(allocator, a, b), arrayList.items[idx], arrayList, target);
}

fn isPossible(idx: usize, a: u64, b: u64, arrayList: *std.ArrayList(u64), target: u64) bool {
    if (idx == arrayList.items.len) {
        return a + b == target or a * b == target;
    }
    return isPossible(idx + 1, a + b, arrayList.items[idx], arrayList, target) or isPossible(idx + 1, a * b, arrayList.items[idx], arrayList, target);
}

pub fn day7_sol(allocator: std.mem.Allocator) !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();

    var part1: u64 = 0;
    var part2: u64 = 0;
    var idx: u32 = 0;

    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(u64))) |line| : (idx += 1) {
        std.debug.print("Processing line {d}\n", .{idx});
        defer allocator.free(line);
        var arrayList = std.ArrayList(u64).init(allocator);
        defer arrayList.deinit();
        var colonIt = std.mem.splitScalar(u8, line, ':');
        var getTarget: bool = true;
        var target: u64 = undefined;
        while (colonIt.next()) |token| {
            if (getTarget) {
                getTarget = false;
                target = try std.fmt.parseInt(u64, token, 10);
                continue;
            }
            var operandIterator = std.mem.splitScalar(u8, std.mem.trim(u8, token, " "), ' ');
            while (operandIterator.next()) |num| {
                try arrayList.append(try std.fmt.parseInt(u64, num, 10));
            }
        }
        if (isPossible(2, arrayList.items[0], arrayList.items[1], &arrayList, target)) {
            part1 += target;
        }
        if (try isPossibleWithConcat(allocator, 2, arrayList.items[0], arrayList.items[1], &arrayList, target)) {
            part2 += target;
        }
    }
    std.debug.print("Part 1 : {d}, Part 2 : {d}\n", .{ part1, part2 });
}

test "day7" {
    try day7_sol(std.testing.allocator);
}
