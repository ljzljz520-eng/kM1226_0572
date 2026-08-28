package input

import (
	"fmt"
	"spacetrash/report"
	"spacetrash/service"
)

type Executor struct {
	Session *service.Session
}

func (e Executor) Execute(command Command) (string, bool, error) {
	switch command.Name {
	case "status":
		return report.RenderSummary(report.BuildSummary(e.Session)), false, nil
	case "move":
		if err := RequireArgs(command, 2); err != nil {
			return "", false, err
		}
		dx, err := IntArg(command, 0)
		if err != nil {
			return "", false, err
		}
		dy, err := IntArg(command, 1)
		if err != nil {
			return "", false, err
		}
		return "", false, e.Session.Move(dx, dy)
	case "scan":
		if err := RequireArgs(command, 1); err != nil {
			return "", false, err
		}
		radius, err := IntArg(command, 0)
		if err != nil {
			return "", false, err
		}
		return fmt.Sprintf("%d contacts", len(e.Session.Scan(radius))), false, nil
	case "fire":
		if err := RequireArgs(command, 2); err != nil {
			return "", false, err
		}
		slot, err := IntArg(command, 0)
		if err != nil {
			return "", false, err
		}
		result, err := e.Session.Fire(slot, command.Args[1])
		if err != nil {
			return "", false, err
		}
		return report.RenderEvent("weapon fired", fmt.Sprintf("damage %d", result.Damage), fmt.Sprintf("destroyed %t", result.Destroyed)), false, nil
	case "upgrade":
		if err := RequireArgs(command, 1); err != nil {
			return "", false, err
		}
		slot, err := IntArg(command, 0)
		if err != nil {
			return "", false, err
		}
		receipt, err := e.Session.UpgradeWeapon(slot)
		if err != nil {
			return "", false, err
		}
		return report.RenderEvent("weapon upgraded", fmt.Sprintf("slot %d", receipt.Slot), fmt.Sprintf("level %d", receipt.NewLevel)), false, nil
	case "turn":
		return "", false, e.Session.AdvanceTurn()
	case "save":
		return "saved", false, e.Session.Save()
	case "help":
		return Help(), false, nil
	case "quit", "exit":
		return "bye", true, nil
	default:
		return "", false, fmt.Errorf("unknown command %s", command.Name)
	}
}
