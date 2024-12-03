const std = @import("std");
const Allocator = std.mem.Allocator;

fn parseMult(data: []u8, length: usize, idxPtr: *u32) u64 {
    var idx: u32 = idxPtr.*;
    defer {
        idxPtr.* = idx;
    }
    while (idx < length) {
        if (idx + 3 >= length or data[idx + 1] != 'u' or data[idx + 2] != 'l' or data[idx + 3] != '(') {
            idx += 1;
            break;
        }
        var num1: u64 = 0;
        var commaFound: bool = false;
        var digitFound: bool = false;
        idx = idx + 4;
        while (idx < length) {
            if (data[idx] == ',') {
                commaFound = true;
                idx += 1;
                break;
            }
            if (data[idx] < '0' or data[idx] > '9') {
                return 0;
            }
            digitFound = true;
            num1 = (num1 * 10) + (data[idx] - '0');
            idx += 1;
        }
        if (!digitFound or !commaFound or num1 > 999) {
            return 0;
        }
        var num2: u64 = 0;
        var bracketFound: bool = false;
        digitFound = false;
        while (idx < length) {
            if (data[idx] == ')') {
                bracketFound = true;
                idx += 1;
                break;
            }
            if (data[idx] < '0' or data[idx] > '9') {
                return 0;
            }
            digitFound = true;
            num2 = (num2 * 10) + (data[idx] - '0');
            idx += 1;
        }
        if (!bracketFound or !digitFound or num2 > 999) {
            return 0;
        }
        return num1 * num2;
    }
    return 0;
}

fn parseDo(data: []u8, length: usize, idxPtr: *u32) ?bool {
    const idx: u32 = idxPtr.*;
    if (idx + 6 <= length and std.mem.eql(u8, data[idx .. idx + 7], "don't()")) {
        return false;
    }
    if (idx + 3 <= length and std.mem.eql(u8, data[idx .. idx + 4], "do()")) {
        return true;
    }
    return null;
}

pub fn day3_sol(allocator: Allocator) !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();
    var part1: u64 = 0;
    var part2: u64 = 0;
    const data = try file.reader().readAllAlloc(allocator, std.math.maxInt(u32));
    defer allocator.free(data);
    const length: usize = data.len;
    var idx: u32 = 0;
    var value: u64 = 0;
    var do: bool = true;
    while (idx < length) {
        if (data[idx] != 'm' and data[idx] != 'd') {
            idx += 1;
            continue;
        }
        if (data[idx] == 'm') {
            value = parseMult(data, length, &idx);
            part1 += value;
            if (do) {
                part2 += value;
            }
        } else {
            const doOrDont: ?bool = parseDo(data, length, &idx);
            if (doOrDont) |val| {
                do = val;
            }
            idx += 1;
        }
    }
    std.debug.print("Part 1 : {d}, Part 2 : {d}\n", .{ part1, part2 });
}
