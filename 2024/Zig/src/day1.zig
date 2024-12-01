const std = @import("std");

pub fn day1_sol() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    var leftArr = std.ArrayList(u32).init(allocator);
    defer leftArr.deinit();
    var rightArr = std.ArrayList(u32).init(allocator);
    defer rightArr.deinit();

    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(usize))) |line| {
        defer allocator.free(line);
        var iterator = std.mem.splitSequence(u8, line, "  ");
        var num1: u32 = undefined;
        if (iterator.next()) |word| {
            num1 = try std.fmt.parseInt(u32, std.mem.trim(u8, word, " "), 10);
        }
        var num2: u32 = undefined;
        if (iterator.next()) |word| {
            num2 = try std.fmt.parseInt(u32, std.mem.trim(u8, word, " "), 10);
        }
        try leftArr.append(num1);
        try rightArr.append(num2);
    }
    const maxRight: u32 = std.mem.max(u32, rightArr.items);
    var hashmap = try allocator.alloc(u32, maxRight + 1);
    @memset(hashmap, @as(u32, 0));
    defer allocator.free(hashmap);

    for (rightArr.items) |number| {
        hashmap[number] += 1;
    }

    std.mem.sort(u32, leftArr.items, {}, comptime std.sort.asc(u32));
    std.mem.sort(u32, rightArr.items, {}, comptime std.sort.asc(u32));

    var part_1_answer: u32 = 0;
    var part_2_answer: u32 = 0;
    for (leftArr.items, rightArr.items) |left, right| {
        if (left > right) {
            part_1_answer += (left - right);
        } else {
            part_1_answer += (right - left);
        }
        //No need to add to hashmap as the frequency of this number in the right list is 0
        if (left >= maxRight) {
            continue;
        }
        part_2_answer += (left * hashmap[left]);
    }
    std.debug.print("Part 1 Answer = {d}, Part 2 Answer = {d}", .{ part_1_answer, part_2_answer });
}
