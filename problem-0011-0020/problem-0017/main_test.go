package main

import "testing"

func TestFindLongestPathToFile(t *testing.T) {
	tables := []struct {
		fs       string
		expected string
	}{
		{"dir\n\tsubdir1\n\tsubdir2\n\t\tsubsubdir1", ""},
		{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", "dir/subdir2/file.ext"},
		{"dir\n\tsubdir1\n\tabcdefghfile.ext\tsubdir2\n\t\tfile.ext", "dir/subdir2/file.ext"},
		{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", "dir/subdir2/subsubdir2/file2.ext"},
		{"file.ext\ndir\n\tsubdir1\n\tsubdir2\n\t\tsubsubdir1", "file.ext"},
		{"file.ext\ndir\n\tsubdir1\n\tsubdir2\n\t\tsubsubdir1\nfile2.ext", "file2.ext"},
	}

	for i, table := range tables {
		actual := findLongestPathToFile(table.fs)
		if actual != table.expected {
			t.Errorf("Test %d: expected %s, actual %s\n", i, table.expected, actual)
		}
	}
}
