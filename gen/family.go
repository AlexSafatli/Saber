package gen

import (
	"github.com/AlexSafatli/Saber/rng"
	"github.com/AlexSafatli/Saber/rpg"
	"sync"
)

const (
	GenderMale   = "Male"
	GenderFemale = "Female"
)

var Genders = []string{GenderMale, GenderFemale}

type Family struct {
	Surname  string
	Origin   *rpg.Region
	Language *Language
}

func GenerateFamily(l *Language, r *rpg.Region) *Family {
	return &Family{
		Surname:  l.Name(),
		Origin:   r,
		Language: l,
	}
}

type FamilyTree struct {
	Root FamilyTreeNode
}

type FamilyTreeNode struct {
	Character   *rpg.Character
	BirthFamily *Family           `json:"-"`
	Mother      *FamilyTreeNode   `json:"-"`
	Father      *FamilyTreeNode   `json:"-"`
	Spouse      *FamilyTreeNode   `json:"-"`
	Children    []*FamilyTreeNode `json:",omitempty"`
}

func (n *FamilyTreeNode) Siblings() []*FamilyTreeNode {
	if n.Father == nil && n.Mother == nil {
		return []*FamilyTreeNode{}
	}
	var parent *FamilyTreeNode
	if n.Father != nil {
		parent = n.Father
	} else {
		parent = n.Mother
	}
	var siblings []*FamilyTreeNode
	for _, c := range parent.Children {
		if c != n {
			siblings = append(siblings, c)
		}
	}
	return siblings
}

func (n *FamilyTreeNode) Married() bool {
	return n.Spouse != nil
}

func (n *FamilyTreeNode) GenerateSpouse(w *rpg.World) {
	var spouseGender string
	if n.Character.Gender == GenderMale {
		spouseGender = GenderFemale
	} else {
		spouseGender = GenderMale
	}
	n.Spouse = generateFamilyTreeNode(
		GenerateFamily(n.BirthFamily.Language,
			&w.Regions[rng.RandomIndex(len(w.Regions))]), spouseGender)
}

func (n *FamilyTreeNode) CanHaveChildren() bool {
	return n.Married() && (n.Character.Gender == GenderFemale ||
		n.Spouse.Character.Gender == GenderFemale)
}

func generateFamilyTreeNode(f *Family, gender string) *FamilyTreeNode {
	node := &FamilyTreeNode{BirthFamily: f,
		Character: GenerateCharacter(f.Language, gender)}
	return node
}

func RandomGender() string {
	return rng.Choose(Genders)
}

func GenerateFamilyTree(f *Family, w *rpg.World,
	numStartingChildren int) *FamilyTree {
	tree := FamilyTree{
		Root: *generateFamilyTreeNode(f, GenderMale),
	}
	var wg sync.WaitGroup
	wg.Add(numStartingChildren)
	tree.Root.GenerateSpouse(w)
	tree.Root.Children = make([]*FamilyTreeNode, numStartingChildren)
	for i := 0; i < numStartingChildren; i++ {
		tree.Root.Children[i] = generateFamilyTreeNode(f, RandomGender())
		tree.Root.Children[i].GenerateSpouse(w)
		go PopulateFamilyTree(tree.Root.Children[i], w, &wg)
	}
	wg.Wait()
	return &tree
}

func PopulateFamilyTree(node *FamilyTreeNode, w *rpg.World,
	wg *sync.WaitGroup) {
	defer wg.Done()
	if !node.CanHaveChildren() {
		return // only populate children if they can have children
	}
	numChildren := rng.RandomSmallNumber()
	wg.Add(numChildren)
	node.Children = make([]*FamilyTreeNode, numChildren)
	for i := 0; i < numChildren; i++ {
		node.Children[i] = generateFamilyTreeNode(node.BirthFamily,
			RandomGender())
		if node.Character.Gender == GenderMale {
			node.Children[i].Father = node
			node.Children[i].Mother = node.Spouse
		} else {
			node.Children[i].Father = node.Spouse
			node.Children[i].Mother = node
		}
		if rng.RandomBoolean() {
			// 50% chance they have a spouse
			node.Children[i].GenerateSpouse(w)
		}
		go PopulateFamilyTree(node.Children[i], w, wg)
	}
}
