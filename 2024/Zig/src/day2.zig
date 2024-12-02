const std = @import("std");

pub fn safe(arr: std.ArrayList(i32)) bool {
    var idx: u32 = 0;
    var inc: bool = true;
    var dec: bool = true;
    var validDiff: bool = true;
    while (idx < arr.items.len - 1) : (idx += 1) {
        if (arr.items[idx] >= arr.items[idx + 1]) {
            inc = false;
        }
        if (arr.items[idx] <= arr.items[idx + 1]) {
            dec = false;
        }
        const diff = arr.items[idx] - arr.items[idx + 1];

        if (@abs(diff) > 3) {
            validDiff = false;
        }
    }
    return (inc or dec) and validDiff;
}

fn getIndicesToRemove(arr: std.ArrayList(i32)) struct { incIdx: u32, decIdx: u32 } {
    var idx: u32 = 0;
    var incIdx: u32 = 0;
    var decIdx: u32 = 0;
    while (idx < arr.items.len - 1) : (idx += 1) {
        if (arr.items[idx] >= arr.items[idx + 1] or arr.items[idx + 1] - arr.items[idx] > 3) {
            incIdx = idx;
            break;
        }
    }
    idx = 0;
    while (idx < arr.items.len - 1) : (idx += 1) {
        if (arr.items[idx] <= arr.items[idx + 1] or arr.items[idx] - arr.items[idx + 1] > 3) {
            decIdx = idx;
            break;
        }
    }
    return .{ .incIdx = incIdx, .decIdx = decIdx };
}

pub fn day2_sol() !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    var part_1_answer: i32 = 0;
    var part_2_answer: i32 = 0;

    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(i32))) |line| {
        defer allocator.free(line);
        var arr = std.ArrayList(i32).init(allocator);
        defer arr.deinit();
        var iterator = std.mem.splitScalar(u8, line, ' ');
        while (iterator.next()) |numberStr| {
            const num = try std.fmt.parseInt(i32, numberStr, 10);
            try arr.append(num);
        }
        if (safe(arr)) {
            part_1_answer += 1;
            part_2_answer += 1;
        } else {
            const indices = getIndicesToRemove(arr);
            var tempArr = std.ArrayList(i32).init(allocator);
            defer tempArr.deinit();
            try tempArr.appendSlice(arr.items[0..indices.incIdx]);
            try tempArr.appendSlice(arr.items[indices.incIdx + 1 ..]);
            if (safe(tempArr)) {
                part_2_answer += 1;
                continue;
            }
            tempArr.clearRetainingCapacity();
            try tempArr.appendSlice(arr.items[0 .. indices.incIdx + 1]);
            try tempArr.appendSlice(arr.items[indices.incIdx + 2 ..]);
            if (safe(tempArr)) {
                part_2_answer += 1;
                continue;
            }
            tempArr.clearRetainingCapacity();
            try tempArr.appendSlice(arr.items[0..indices.decIdx]);
            try tempArr.appendSlice(arr.items[indices.decIdx + 1 ..]);
            if (safe(tempArr)) {
                part_2_answer += 1;
                continue;
            }
            tempArr.clearRetainingCapacity();
            try tempArr.appendSlice(arr.items[0 .. indices.decIdx + 1]);
            try tempArr.appendSlice(arr.items[indices.decIdx + 2 ..]);
            if (safe(tempArr)) {
                part_2_answer += 1;
                continue;
            }
        }
    }
    std.debug.print("Part 1 : {d}, Part 2 : {d}\n", .{ part_1_answer, part_2_answer });
}
