package cmd

import (
	"filmPrice/config"
	filmDao "filmPrice/internal/apps/film/dao"
	systemDao "filmPrice/internal/apps/system/dao"
	taskDao "filmPrice/internal/apps/task/dao"
	"filmPrice/pkg/logger"
	"filmPrice/pkg/password"
	"github.com/spf13/cobra"
)

var filePath2 string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化项目",
	Long:  `初始化项目`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config.LoadConfigYaml(filePath2)

		if err := config.Get().Mysql.Init(logger.New(logger.WithDefault(), logger.WithPrefix("Init "))); err != nil {
			return err
		}

		// 初始化数据库
		config.GetDB().AutoMigrate(
			&systemDao.SystemUserModel{},
			&systemDao.SystemGroupModel{},
			&systemDao.SystemPermModel{},
			&taskDao.TaskModel{},
			&taskDao.TaskLogModel{},
			&filmDao.FilmModel{},
			&filmDao.FilmLinkModel{},
			&filmDao.FilmPriceModel{},
			&filmDao.FilmPriceHistoryModel{},
		)

		// 创建管理用户
		pwd, _ := password.GenPassword("admin")
		config.GetDB().FirstOrCreate(&systemDao.SystemUserModel{Name: "admin", Password: pwd, Type: "local"})

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	startCmd.Flags().StringVarP(&filePath2, "conf", "c", "config.yaml", "配置文件路径")
}
