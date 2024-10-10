package types

type AuthenticationRequest struct {
	Username           *string                `json:"username,omitempty"`
	Password           *string                `json:"password,omitempty"`
	AuthenticationMode MstrAuthenticationMode `json:"loginMode"`
	ApplicationType    *MstrApplicationType   `json:"applicationType,omitempty"`
}

type MstrApplicationType int

const (
	BulkTranslationTool       MstrApplicationType = 30
	CacheUtility              MstrApplicationType = 19
	COMBrowser                MstrApplicationType = 28
	CommandLineInterface      MstrApplicationType = 15
	CommandManager            MstrApplicationType = 13
	ConfigurationWizard       MstrApplicationType = 17
	CubeAdvisor               MstrApplicationType = 32
	CustomApp                 MstrApplicationType = 8
	DossierMobile             MstrApplicationType = 36
	DossierWeb                MstrApplicationType = 35
	DSSScheduler              MstrApplicationType = 7
	DSSWeb                    MstrApplicationType = 6
	EnterpriseManager         MstrApplicationType = 14
	FireEvent                 MstrApplicationType = 20
	HealthCenter              MstrApplicationType = 31
	HyperBrowserChrome        MstrApplicationType = 48
	HyperDesktopMac           MstrApplicationType = 58
	HyperDesktopWindows       MstrApplicationType = 59
	HyperMessagingSlack       MstrApplicationType = 57
	HyperMobileAndroid        MstrApplicationType = 50
	HyperMobileIOS            MstrApplicationType = 49
	HyperOfficeOutlookAndroid MstrApplicationType = 55
	HyperOfficeOutlookIOS     MstrApplicationType = 54
	HyperOfficeOutlookMac     MstrApplicationType = 53
	HyperOfficeOutlookWeb     MstrApplicationType = 51
	HyperOfficeOutlookWindows MstrApplicationType = 52
	HyperScreenAndroidTV      MstrApplicationType = 63
	HyperScreenAppleTV        MstrApplicationType = 61
	HyperScreenFireTV         MstrApplicationType = 62
	HyperSDK                  MstrApplicationType = 60
	HyperVoiceAlexa           MstrApplicationType = 56
	Jupyter                   MstrApplicationType = 64
	LibraryMobileAndroid      MstrApplicationType = 39
	MacWorkstation            MstrApplicationType = 37
	MDScan                    MstrApplicationType = 18
	MDUpdate                  MstrApplicationType = 27
	MicrosoftOffice           MstrApplicationType = 47
	Mobile                    MstrApplicationType = 29
	NarrowcastServer          MstrApplicationType = 9
	ObjectManager             MstrApplicationType = 10
	ODBOCubeDesigner          MstrApplicationType = 12
	ODBOProvider              MstrApplicationType = 11
	Office                    MstrApplicationType = 1
	OfficeWS                  MstrApplicationType = 23
	OneTier                   MstrApplicationType = 34
	Portal                    MstrApplicationType = 25
	PowerBI                   MstrApplicationType = 46
	Product01                 MstrApplicationType = 66
	Product02                 MstrApplicationType = 67
	Product03                 MstrApplicationType = 68
	Product04                 MstrApplicationType = 69
	Product05                 MstrApplicationType = 70
	Product06                 MstrApplicationType = 71
	ProjectBuilder            MstrApplicationType = 16
	ProjectUpgrade            MstrApplicationType = 5
	Qlik                      MstrApplicationType = 45
	RStudio                   MstrApplicationType = 65
	Server                    MstrApplicationType = 4
	ServerAdmin               MstrApplicationType = 2
	Tableau                   MstrApplicationType = 44
	Tester                    MstrApplicationType = 26
	Tools                     MstrApplicationType = 24
	TypeJavaDesktop           MstrApplicationType = 21
	WebAdmin                  MstrApplicationType = 3
	WebServices               MstrApplicationType = 22
	WindowsWorkstation        MstrApplicationType = 37
	Workstation               MstrApplicationType = 37
	WorkstationMigration      MstrApplicationType = 72
)
