package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-item-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		MaintenanceItem:              data.MaintenanceItem,
		MaintenanceItemDescription:   data.MaintenanceItemDescription,
		MaintenanceStrategy:          data.MaintenanceStrategy,
		MaintenancePlanCategory:      data.MaintenancePlanCategory,
		MaintenancePlanCallObject:    data.MaintenancePlanCallObject,
		MaintenancePlan:              data.MaintenancePlan,
		MaintenancePlanItemPosition:  data.MaintenancePlanItemPosition,
		MaintenanceItemObjectList:    data.MaintenanceItemObjectList,
		FunctionalLocationLabelName:  data.FunctionalLocationLabelName,
		Equipment:                    data.Equipment,
		Assembly:                     data.Assembly,
		AdditionalDeviceData:         data.AdditionalDeviceData,
		TaskListType:                 data.TaskListType,
		TaskListGroup:                data.TaskListGroup,
		TaskListGroupCounter:         data.TaskListGroupCounter,
		OperationSystemCondition:     data.OperationSystemCondition,
		NumberOfTaskListExecutions:   data.NumberOfTaskListExecutions,
		MaintNotifTskIsAutomlyDtmnd:  data.MaintNotifTskIsAutomlyDtmnd,
		MaintenancePlanningPlant:     data.MaintenancePlanningPlant,
		MaintenancePlannerGroup:      data.MaintenancePlannerGroup,
		MaintenanceOrderType:         data.MaintenanceOrderType,
		NotificationType:             data.NotificationType,
		MaintenanceActivityType:      data.MaintenanceActivityType,
		MainWorkCenter:               data.MainWorkCenter,
		MainWorkCenterPlant:          data.MainWorkCenterPlant,
		MaintPriority:                data.MaintPriority,
		MaintPriorityType:            data.MaintPriorityType,
		BusinessArea:                 data.BusinessArea,
		ImmediateReleaseIsBlocked:    data.ImmediateReleaseIsBlocked,
		Material:                     data.Material,
		SerialNumber:                 data.SerialNumber,
		ServiceDocumentType:          data.ServiceDocumentType,
		ServiceContract:              data.ServiceContract,
		ServiceContractItem:          data.ServiceContractItem,
		ServiceOrderTemplate:         data.ServiceOrderTemplate,
		ServiceDocumentPriority:      data.ServiceDocumentPriority,
		Product:                      data.Product,
		MaintenancePlant:             data.MaintenancePlant,
		AssetLocation:                data.AssetLocation,
		AssetRoom:                    data.AssetRoom,
		PlantSection:                 data.PlantSection,
		WorkCenter:                   data.WorkCenter,
		ABCIndicator:                 data.ABCIndicator,
		MaintObjectFreeDefinedAttrib: data.MaintObjectFreeDefinedAttrib,
		CompanyCode:                  data.CompanyCode,
		MasterFixedAsset:             data.MasterFixedAsset,
		FixedAsset:                   data.FixedAsset,
		LocAcctAssgmtBusinessArea:    data.LocAcctAssgmtBusinessArea,
		CostCenter:                   data.CostCenter,
		ControllingArea:              data.ControllingArea,
		WBSElement:                   data.WBSElement,
		SettlementOrder:              data.CycleSetSequence,
		CycleSetSequence:             data.CycleSetSequence,
		StandingOrderNumber:          data.StandingOrderNumber,
		CreationDate:                 data.CreationDate,
		LastChangeDate:               data.LastChangeDate,
		LastChangeDateTime:           data.LastChangeDateTime,
		// ToCallObjects:                           data.ToMaintPlanCallObjects.ToMaintPlanCallObjectsResults[0],
	}

	return item, nil
}
