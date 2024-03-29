/*
 * Control-M Services
 *
 * Provides access to BMC Control-M Services
 *
 * API version: 9.20.115
 * Contact: customer_support@bmc.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package controlm_go

type SlaService struct {
	// Service Name
	ServiceName string `json:"serviceName,omitempty"`
	// Status Code
	Status string `json:"status,omitempty"`
	// Status Reason
	StatusReason string `json:"statusReason,omitempty"`
	// Start Time
	StartTime string `json:"startTime,omitempty"`
	// End Time
	EndTime string `json:"endTime,omitempty"`
	// Due Time
	DueTime string `json:"dueTime,omitempty"`
	// Slack Time
	SlackTime string `json:"slackTime,omitempty"`
	// Service Order DateTime
	ServiceOrderDateTime string `json:"serviceOrderDateTime,omitempty"`
	// Control-M Order Date
	ScheduledOrderDate string `json:"scheduledOrderDate,omitempty"`
	// Service Key
	ServiceJob string `json:"serviceJob,omitempty"`
	// Service Control-M
	ServiceControlM string `json:"serviceControlM,omitempty"`
	// Priority
	Priority string `json:"priority,omitempty"`
	// User note
	Note string `json:"note,omitempty"`
	// Number of Jobs
	TotalJobs string `json:"totalJobs,omitempty"`
	// Jobs Completed
	JobsCompleted string `json:"jobsCompleted,omitempty"`
	// Jobs without statistics
	JobsWithoutStatistics string `json:"jobsWithoutStatistics,omitempty"`
	// Completion Percentage
	CompletionPercentage string `json:"completionPercentage,omitempty"`
	// Average Completion Time
	AverageCompletionTime string `json:"averageCompletionTime,omitempty"`
	// Error details
	Errors string `json:"errors,omitempty"`
	StatusByJobs *SlaServiceStatusByJobs `json:"statusByJobs,omitempty"`
}
