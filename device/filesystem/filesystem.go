package filesystem

const Separator = "/"

type Item interface {
	TotalSize() int64
}

type File struct {
	Name string
	Size int64
}

func (f File) TotalSize() int64 {
	return f.Size
}

type Dir struct {
	Name     string
	parent   *Dir
	children []*Dir
	files    []*File
}

func (d *Dir) AddDir(c *Dir) {
	c.parent = d
	d.children = append(d.children, c)
}

func (d *Dir) AddFile(c *File) {
	d.files = append(d.files, c)
}

func (d *Dir) Contents() []Item {
	contents := make([]Item, len(d.children)+len(d.files))
	i := 0

	for _, child := range d.children {
		contents[i] = child
		i += 1
	}

	for _, file := range d.files {
		contents[i] = file
		i += 1
	}

	return contents
}

func (d *Dir) GetDir(name string) *Dir {
	for _, aDir := range d.children {
		if aDir.Name == name {
			return aDir
		}
	}
	return nil
}

func (d *Dir) Parent() *Dir {
	return d.parent
}

func (d *Dir) TotalSize() int64 {
	totalSize := int64(0)
	for _, item := range d.Contents() {
		totalSize += item.TotalSize()
	}
	return totalSize
}

type FS struct {
	root *Dir
}

func NewFS() *FS {
	root := Dir{}
	return &FS{root: &root}
}

func (fs *FS) Root() *Dir {
	return fs.root
}

func (fs *FS) Dirs() []*Dir {
	var dirs []*Dir
	return addDir(fs.root, dirs)
}

func addDir(from *Dir, to []*Dir) []*Dir {
	to = append(to, from)
	for _, c := range from.children {
		to = addDir(c, to)
	}
	return to
}
