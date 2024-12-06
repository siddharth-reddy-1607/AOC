const std = @import("std");

fn customSort(context: *std.AutoHashMap(u32, std.AutoHashMap(u32, bool)), x: u32, y: u32) bool {
    const valMap = context.get(x) orelse null;
    if (valMap == null) {
        return false;
    }
    return valMap.?.get(y) orelse false;
}

fn isValidOrder(array: *std.ArrayList(u32), hashmap: *std.AutoHashMap(u32, std.AutoHashMap(u32, bool))) bool {
    for (0..array.items.len) |i| {
        for (i + 1..array.items.len) |j| {
            const valMap = hashmap.get(array.items[i]) orelse null;
            if (valMap == null) {
                return false;
            }
            const found = valMap.?.get(array.items[j]) orelse false;
            if (!found) {
                return false;
            }
        }
    }
    return true;
}

fn makeOrderValid(array: *std.ArrayList(u32), hashmap: *std.AutoHashMap(u32, std.AutoHashMap(u32, bool))) void {
    std.mem.sort(u32, array.items, hashmap, customSort);
}

pub fn day5_sol(allocator: std.mem.Allocator) !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();

    var part1: u32 = 0;
    var part2: u32 = 0;
    var hashmap = std.AutoHashMap(u32, std.AutoHashMap(u32, bool)).init(allocator);
    defer {
        var valueIt = hashmap.valueIterator();
        while (valueIt.next()) |valueMap| {
            valueMap.*.deinit();
        }
        hashmap.deinit();
    }
    var getOrderings = false;
    while (try file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(usize))) |line| {
        defer allocator.free(line);
        var array = std.ArrayList(u32).init(allocator);
        defer array.deinit();
        if (std.mem.eql(u8, line, "")) {
            getOrderings = true;
            continue;
        }
        if (getOrderings) {
            var it = std.mem.splitScalar(u8, line, ',');
            while (it.next()) |numStr| {
                try array.append(try std.fmt.parseInt(u8, numStr, 10));
            }
            if (isValidOrder(&array, &hashmap)) {
                part1 += array.items[array.items.len / 2];
            } else {
                makeOrderValid(&array, &hashmap);
                part2 += array.items[array.items.len / 2];
            }
        } else {
            var it = std.mem.splitScalar(u8, line, '|');
            var edge = [2]u32{ 0, 0 };
            var idx: u32 = 0;
            while (it.next()) |numStr| : (idx = idx + 1) {
                edge[idx] = try std.fmt.parseInt(u8, numStr, 10);
            }
            const hashmapGOP = try hashmap.getOrPut(edge[0]);
            if (hashmapGOP.found_existing) {
                try hashmapGOP.value_ptr.*.put(edge[1], true);
            } else {
                hashmapGOP.value_ptr.* = std.AutoHashMap(u32, bool).init(allocator);
                try hashmapGOP.value_ptr.*.put(edge[1], true);
            }
        }
    }
    std.debug.print("Part 1 : {d},Part 2 : {d}\n", .{ part1, part2 });
}
