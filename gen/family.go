package gen

import (
	"../entities"
)

type Family struct {
	Locale   *entities.Region
	Language *Language
}

type FamilyTree struct {
	Root FamilyTreeNode
}

type FamilyTreeNode struct {
	Character   *entities.Character
	BirthFamily *Family
	Mother      *FamilyTreeNode
	Father      *FamilyTreeNode
	Spouse      *FamilyTreeNode
	Children    []*FamilyTreeNode
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

func (n *FamilyTreeNode) GenerateSpouse() {
	//var spouseGender string
	//if n.Character.Gender == "Male" {
	//  spouseGender = "Female"
	//} else {
	//  spouseGender = "Male"
	//}
}

func (n *FamilyTreeNode) CanHaveChildren() bool {
	return n.Married() && (n.Character.Gender == "Female" || n.Spouse.Character.Gender == "Female")
}

func generateFamilyTreeNode(f *Family, gender string) *FamilyTreeNode {
	node := &FamilyTreeNode{BirthFamily: f, Character: GenerateCharacter(f.Language, gender)}
	return node
}
