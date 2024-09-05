package checklist

import (
	"github.com/gofiber/fiber/v2"
)

type ChecklistHandler struct {
	repo *ChecklistRepository
}

func NewChecklistHandler(repo *ChecklistRepository) *ChecklistHandler {
	return &ChecklistHandler{repo: repo}
}

func (h *ChecklistHandler) CreateChecklist(c *fiber.Ctx) error {
	var request struct {
		Title string `json:"title"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	err := h.repo.CreateChecklist(request.Title)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create checklist"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "checklist created"})
}

func (h *ChecklistHandler) GetChecklists(c *fiber.Ctx) error {
	checklists, err := h.repo.GetChecklists()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch checklists"})
	}

	return c.JSON(checklists)
}

func (h *ChecklistHandler) DeleteChecklist(c *fiber.Ctx) error {
	checkListID := c.Params("id")

	err := h.repo.DeleteChecklist(checkListID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not delete checklist"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "checklist deleted"})
}

func (h *ChecklistHandler) GetItems(c *fiber.Ctx) error {
	checkListID := c.Params("id")

	checkListItem, err := h.repo.GetCheckListItem(checkListID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch checklist item"})
	}

	return c.JSON(checkListItem)
}

func (h *ChecklistHandler) AddItem(c *fiber.Ctx) error {
	checklistID := c.Params("id")

	var request struct {
		Item string `json:"item"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	err := h.repo.AddItem(checklistID, request.Item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not add item to checklist"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "item added"})
}

func (h *ChecklistHandler) UpdateItem(c *fiber.Ctx) error {
	checkListID := c.Params("id")
	itemID := c.Params("idItem")

	var request struct {
		Item string `json:"item"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	err := h.repo.UpdateItem(checkListID, itemID, request.Item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update item to checklist"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "item updated"})
}

func (h *ChecklistHandler) DeleteItem(c *fiber.Ctx) error {
	checkListID := c.Params("id")
	itemID := c.Params("idItem")

	err := h.repo.DeleteItem(checkListID, itemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not delete item to checklist"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "item deleted"})
}

func (h *ChecklistHandler) UpdateItemStatus(c *fiber.Ctx) error {
	checkListID := c.Params("id")
	itemID := c.Params("idItem")

	var request struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	err := h.repo.UpdateItemStatus(checkListID, itemID, request.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update item to checklist"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "item status updated"})
}
