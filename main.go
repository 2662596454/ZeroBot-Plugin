// 资源ZeroBot-Plugin主文件
包包主

进口进口（
	“编码/json” “编码/json”
	“旗帜” “旗帜”
	“FMMT” “FMMT”
	《数学/兰德》 《数学/兰德》
	“操作系统”
	“运行” “运行时”
	“strcon” “strconv”
	“弦”“弦”
	“时间”“时间”

	_ "github.com/FloatTech/ZeroBot-Plugin/console"  // 更改控制台属性"github.com/FloatTech/ZeroBot-Plugin/console" // 更改控制台属性"github.com/FloatTech/ZeroBot-Plugin/console " // 更改控制台属性"github.com/FloatTech /ZeroBot-Plugin/console" // 更改控制台属性 console" // 更改控制台属性"github.com/FloatTech /ZeroBot-Plugin/console" // 更改控制台属性"github.com/FloatTech/ZeroBot-Plugin/console" // 更改控制台属性"github.com/FloatTech/ZeroBot-Plugin/console" // 更改控制台属性"github.com/FloatTech/ZeroBot-Plugin/console" // 更改控制台属性"github.com/FloatTech /ZeroBot-Plugin/console" // 更改控制台属性

	“github.com/FloatTech/ZeroBot-Plugin/看板”//打印横幅"github.com/FloatTech/ZeroBot-Plugin/kanban" // 打印横幅 // 打印横幅"github.com/FloatTech/ZeroBot-Plugin/kanban " // 打印横幅// 打印横幅"github.com/FloatTech/ZeroBot-Plugin/kanban" // 打印横幅 // 打印横幅"github.com/FloatTech/ZeroBot-Plugin/kanban " // 打印横幅

	// ---------以下插件通过前面加 // 注释、注释后没有加载插件-------- //// ---------以下插件欢迎通过前面添加 // 注释、注释后不再加载--------- //
	// ---------------------- 插件优先级按顺序从高到低------------------ - - - - - - ------ //// ---------------------- 插件优先级按顺序从高到低--- - - ---- - - - --- ------------ //----- //// ------------------ ---- --- 插件优先级按顺序从高到低----- ---- - - - --- ------------ //// --- ---- --------------- 插件优先级按顺序从高到低---------------------------- ------ - - - - ----- //// ----------------- ----- 插件优先级按顺序从高到低---------- - - - --- ------------ //// ------- --- 插件优先级按顺序从高到低------------------ - - - - - - ----- //// ----- ----------------- 插件优先级按顺序从高到低---- - ---- - - - --- --------- --- //----- //// ----------------  ----- //// ---------------------- 插件优先级按顺序从高到低--- - - ---- - - - --- ------------ //----- //// ------------------ ------- 插件优先级按顺序从高到低----- ---- - - - --- ------------ //// --- ------------------- 插件优先级按顺序从高到低-------------------------------- - - - - ----- //// ----------------- ----- 插件优先级按顺序从高到低--------- - - - --- ------------ //// ---------------------- 插件优先级按顺序从高到低------------------ - - - - - - ----- //// ---------------------- 插件优先级按顺序从高到低---- - ---- - - - --- ------------ //----- //// ---------------------- 插件优先级按顺序从高到低----- ---- - - - --- ------------ //// ---------------------- 插件优先级别按顺序从高到低------------------ - - - - ----- //// ------------ ----- ----- 插件优先级按顺序从高到低--------- - - - --- ------------ //
	// //// // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
	// //// // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
	// //// // // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // ///// //// // // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
	// //// // // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //
	// //// // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // //// // // // // // // // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // // // // // // // // // // // // // //
	// ---------------------------------------- 高优先级区---- ------------ ------------ //// ---------------------- ------高级优先级区---------------------------------------- ------------------------------------------------ //// ---------- ------------------------ ------ 高优先级区---- ------------ ------------ //// -------- -------------- ------ 高级优先级区-- ------------------------ -- -------------- ---------- -- //// ------------------ ---------------------- 高级优先级区---- ------------ -------- ---- //// ---------------------------------- ------ 高级优先级区---------- ------------------------------------------ -------------------- ---------------- //// ------------------------------------------ ---- ------ 高优先级区---- ------------ ------------ //// ---- ---- -------------- -------------- //// ---------- ------------------------ ------ 高优先级区---- ------------ ------------ //// -------- -------------- ------ 高级优先级区-- -------------------------- -------------- ---------- -- //// ---------------------------------------- 高优先级区---- ------------ ------------ //// ---------------------- ------高级优先级区---------------------------------------- ------------------------------------ //// ---------------------------------- ------ 高优先级区---- ------------ ------------ //// -------- -------------- ------ 高级优先级区---------------------------- -------------- ------------ //
	// vvvvvvvvvvvv 高级区vvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv 高级区vvvvvvvvvvvv vvvviv i //// v vvvvvvvvvvv vv 高级区vvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv高级区vvvvvvvvvvvvv vvvvVivi //// vvvvvvvvvvvvv 高级区vvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv 高级区vvvvvvvvvvvv vvvvviv i //// v vvvvvvvvvvvv vv高级区vvvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv高级区vvvvvvvvvvvvv vvvvVivi //
	// vvvvvvvvvvvv 高级区vvvvvvvvv //// vvvvvvvvvvvv 高级区vvvvvvvvvv //// vvvvvvvvvvvv 高级区vvvvvvvvvv //// vvvvvvvvvvv 高级区vvvvvvvvvv //// vvvvvvvvvvvv 高级区vvvvvvvvv //// vvvvvvvvvvvv 高级区vvvvvvvvvvv //// vvvvvvvvvvvv 高级区vvvvvvvvvv //// vvvvvvvvvvv 高级区vvvvvvvvvvv //
	// vvvvvvv 高级区vvvvvvv //// vvvvvvv 高级区vvvvvvv //
	// vvvvvvvvvvvv //// vvvvvvvvvvvv //// vvvvvvvvvvvvv //// vvvvvvvvvvvv //
	// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //

	_“github . com/FloatTech/ZeroBot-Plugin/plugin/antiabuse”. com/FloatTech/ZeroBot-Plugin/plugin/antiabuse”

	_“github .com/FloatTech/ZeroBot-Plugin/plugin/chat” // 基础词库"github.com/FloatTech/ZeroBot-Plugin/plugin/chat" // 基础词库// 基础词库"github.com/FloatTech/ZeroBot-Plugin/plugin/chat" // 基础词库

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/sleepmanage" // 统计睡眠时间"github.com/FloatTech/ZeroBot-Plugin/plugin/sleepmanage" // 统计睡眠时间"github.com/FloatTech/ZeroBot-Plugin/plugin/sleepmanage" // 统计睡眠时间"github.com/FloatTech/ZeroBot-Plugin/plugin/sleepmanage" // 统计睡眠时间

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/atri" // ATRI词库"github.com/FloatTech/ZeroBot-Plugin/plugin/atri" // ATRI词库"github.com/FloatTech/ZeroBot- Plugin/plugin/atri" // ATRI词库"github.com/ FloatTech/ZeroBot-Plugin/plugin/atri" // ATRI词库

	_“github . com/FloatTech/ZeroBot-Plugin/plugin/manager”   // 群管 "github.com/FloatTech/ZeroBot-Plugin/plugin/manager" // 群管"github.com/FloatTech/ZeroBot-Plugin/plugin/manager" // 群管 "github.com/FloatTech/ZeroBot-Plugin/plugin/manager" // 群管

	_ "github.com/FloatTech/zbputils/job" // 定时指令触发"github.com/FloatTech/zbputils/job" // 定时指令触发"github.com/FloatTech/zbputils/job" // 定时指令触发" github.com/FloatTech/zbputils/job" //定时触发指令"github.com/FloatTech/zbputils/job"    // 定时指令触发"github.com/FloatTech/zbputils/job" // 定时指令触发"github.com/FloatTech/zbputils/job" // 定时指令触发" github .com/FloatTech/zbputils/job" //定时触发指令

	// ^^^^ //// ^^^^ //// ^^^^ //// ^^^^ //
	// ^^^^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^^^^ //// ^^^^^ ^^^ ^^ ^^^^^^^^^^ //// ^^^^^^^^^^^^^^^^^ //
	// ^^^^^^^高级区^^^^^^^ //// ^^^^^^^高级区^^^^^^^ //// ^^^^^^^^^ ^ ^高级^区^^^^^^ //// ^^^^^^^^高级区^^^^^^^ //// ^^^^^^^高级区^^^^^^^ //// ^^^^^^^高级区^^^^^^^ //// ^^^^^^^^^ ^ ^高级^区^^^^^^ //// ^^^^^^^^高级区^^^^^^^ //
	// ^^^^^^^^^^^^^^^^^^^^^^高级区^^^^^^^^^^^^^^^^^^^^^ //// ^^ ^^ ^^ ^^^ ^ ^^^ ^^^^高级区^^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^ ^^^ ^^^ ^高级区^^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^ ^^^ ^^^^高级区^^^ ^^^^^ ^^^^^^^^ //^^高级区^^^^^^^^^^^^^^^^^^^^^ //// ^^ ^^ ^^ ^^^ ^ ^^^ ^^^^高级区^^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^ ^^^ ^^^ ^高级区^^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^ ^^^ ^^^^高级区^^^ ^^^^^ ^^^^^^^^ //// ^^^^^^^^^^^^^^^^^^^^高级区^^^^^^^^^^^^^^^^^^^^ //// ^^^^ ^^ ^^^ ^ ^^^ ^^^^高级区^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^^^ ^^^ ^高级区^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^ ^^^ ^^^^高级区^^^ ^^^^^ ^^^^^^^^ //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^高级区^^ ^^^ ^^^^^ ^^^^ ^^^^ ^^^ ^^^^^^^^^^^^^^^ //// ^^^^^^ ^^^^^^ ^^^ ^^^^^^ ^^^^ ^^^^ ^^^高级区^^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^ ^^^^ ^^^^^^ /// / ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^^高级区^^^^^^^^^^^^^^^ ^^^ ^^^ ^^^^^^^^^^^^^^^^^^^ //// ^ ^^^ ^^^^ ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^高级区^^^^^^^^^^^^^^^ ^^^^^^^^^^^^^^^^^^^^^^^^^^^ //// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^高级区^^^^^ ^^^^^ ^^^^ ^^^^ ^^^ ^^^^^^^^^^^^^^^ //// ^^^^^^ ^^^^^^^^^ ^^^^^^ ^^^^ ^^^^ ^^^高级区^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^^^^^ ^^^^^^ /// / ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^高级区^^^^^^^^^^^^^^ ^^^ ^^^ ^^^^^^^^^^^^^^^^^^ //// ^^^^ ^^^^ ^^^^^^^^^^^^^^^^^^^^^^^^^^^高级区^^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^^^^^^^^^^^ //
	// ---------------------------------------- 高优先级区---- ------------ ------------ //// ---------------------- ------高级优先级区---------------------------------------- ------------------------ //// ---------------------- ------------ ------ 高优先级区---- ------------ ------------ //// -------- -------------- ------ 高级优先级区---------------- -- ------------ -------------- ------------ // ------------ -------------- ------------ //// ---------------------------------------- 高优先级区---- ------------ ------------ //// ---------------------- ------高级优先级区---------------------------------------- ------------------------ //// ---------------------- ------------ ------ 高优先级区---- ------------ ------------ //// -------- -------------- ------ 高级优先级区---------------- ------------ -------------- ------------ //
	// //// // // // // // // // // // // // // // // //// //// // // // // // // // // // // // // // //
	// //// // // // // // // // // // // // // // //// //// // // // // // // //
	// //// // // // // // // // //// //// // // // // // // //// //// // // // // // // //// //// // // // // // //
	// //// // // // // // // // //// //// // // // // // // //// //// // // // // // // //// //// // // // // // //
	// //// // // // // // // //// //// // // // // // //// //// // // // // // //// //// // // // // //
	// ----------------------------------------中优先级区---- ------------ ------------ //// ---------------------- ------中优先级区---------------------------- //// -------- --------------------------------中优先级区---- ---------- -- ------------ //// ---------------------- ------中优先级区---------------------------- //// ----------------------------------------中优先级区---- ------------ ------------ //// ---------------------- ------中优先级区---------------------------- //// -------- --------------------------------中优先级区---- ---------- -- ------------ //// ---------------------- ------中优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvv vv维维//// vvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvv呜呜呜//// vvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvv vv维维//// vvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvvvvv //// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvv呜呜呜//
	// vvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvv //
	// vvvvvvv中优先级区vvvvvvv //// vvvvvv中优先级区vvvvvvv //// vvvvvvv中优先级区vvvvvvv //// vvvvvvv中优先级区vvvvvvv //
	// vvvvvvvvvvvvv //// vvvvvvvvvvvvv //// vvvvvvvvvvvvvv //// vvvvvvvvvvvvv //// vvvvvvvvvvvvv //// vvvvvvvvvvvvv //// vvvvvvvvvvvvvv //// vvvvvvvvvvvvv //
	// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //// 呜呜呜 //

// ahsai tts"github.com/FloatTech/ZeroBot-Plugin/plugin/ahsai" // ahsai tts"github.com/FloatTech/ZeroBot-Plugin/插件/ahsai" // ahsai tts"github.com/FloatTech/ZeroBot -Plugin/plugin/ahsai" // ahsai tts //_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ahsai" // ahsai tts"github.com/FloatTech/ZeroBot-Plugin/plugin/ahsai" // ahsai tts"github.com/FloatTech/ZeroBot-Plugin/plugin/ahsai" // ahsai tts"github.com/FloatTech/ZeroBot-Plugin/plugin/ahsai" // ahsai tts
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse" // 服务器监控"github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse" // 服务器监控"github.com/FloatTech/ZeroBot-Plugin/ plugin/aifalse" // 服务器监控"github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse" // 服务器监控"github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse"  // 服务器监控"github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse" // 服务器监控"github.com/FloatTech/ZeroBot-Plugin/plugin /aifalse" // 监控服务器"github.com/FloatTech/ZeroBot-Plugin/plugin/aifalse" // 服务器监控
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/aipaint"  // ai 绘图"github.com/FloatTech/ZeroBot-Plugin/plugin/aipaint" // ai 绘图"github.com/FloatTech/ZeroBot-Plugin/plugin /aipaint" // ai 绘图"github.com/FloatTech/ZeroBot-Plugin/plugin/aipaint" // ai 绘图
	_“github.com/FloatTech/ZeroBot-Plugin/plugin/aiwife”.com/FloatTech/ZeroBot-Plugin/plugin/aiwife” 
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/alipayvoice" 
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw"  // 触发者撤回时也自动撤回"github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回"github .com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回"github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回 github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回"github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回"github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回"github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回"github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw" // 触发者撤回时也自动撤回"github.com/FloatTech/ZeroBot-Plugin/plugin/autowithdraw"     // 触发者撤回时也自动撤回
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/baidu" // 百度一下"github.com/FloatTech/ZeroBot-Plugin/plugin/baidu" // 百度一下"github.com/FloatTech/ZeroBot-Plugin/plugin/baidu" // 百度一下"github.com/FloatTech/ZeroBot-Plugin/plugin/baidu"            // 百度一下
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/baiduaudit" // 百度内容审核"github.com/FloatTech/ZeroBot-Plugin/plugin/baiduaudit" // 百度内容审核"github.com/FloatTech/ZeroBot- Plugin/plugin/baiduaudit" // 百度内容审核"github.com/FloatTech/ZeroBot-Plugin/plugin/baiduaudit" // 百度内容审核"github.com/FloatTech/ZeroBot-Plugin/plugin/baiduaudit"  // 百度内容审核"github.com/FloatTech/ZeroBot-Plugin/plugin/baiduaudit" // 百度内容审核"github.com/FloatTech/ZeroBot-Plugin /plugin/baiduaudit" // 百度内容审核"github.com/FloatTech/ZeroBot-Plugin/plugin/baiduaudit" // 百度内容审核
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/base16384” 
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/base64gua" //base64卦加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/base64gua" //base64卦加解密"github.com/FloatTech/ ZeroBot-Plugin/plugin/base64gua" // base64卦加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/base64gua" // base64卦加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/base64gua"  //base64卦加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/base64gua" //base64卦加解密"github.com/FloatTech/ZeroBot -Plugin/plugin/base64gua" //base64卦加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/base64gua" //base64卦加解密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro"  //base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro" //base天城文加解密"github.com/FloatTech /ZeroBot-Plugin/plugin/baseamasiro" //base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro" //base天城文加解密base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro" //base天城文加解密"github.com/ FloatTech/ZeroBot-Plugin/plugin/baseamasiro" // base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro" // base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro" // base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro" // base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro" // base天城文加解密"github.com/FloatTech/ZeroBot-Plugin/plugin/baseamasiro"      // base天城文加解密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili"   // b站相关
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview"  // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github .com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录 github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview" // 哀伤雪刃吧推书记录"github.com/FloatTech/ZeroBot-Plugin/plugin/bookreview"       // 哀伤雪刃吧推书记录
	_“github.com/FloatTech/ZeroBot-Plugin/plugin/cangtoushi”"github.com/FloatTech/ZeroBot-Plugin/plugin/cangtoushi" 
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chess" // 国际象棋"github.com/FloatTech/ZeroBot-Plugin/plugin/chess" // 国际象棋"github.com/FloatTech/ZeroBot-Plugin/plugin/chess" // 国际象棋"github.com/FloatTech/ZeroBot-Plugin/plugin/chess"            // 国际象棋
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/choose" // 选择困难症帮手"github.com/FloatTech/ZeroBot-Plugin/plugin/choose" // 选择困难症帮手"github.com/FloatTech/ZeroBot-Plugin/plugin/choose" // 选择困难症帮手"github.com/FloatTech/ZeroBot-Plugin/plugin/choose"           // 选择困难症帮手
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/choufanghua" 
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chrev"  // 中文字符概述"github.com/FloatTech/ZeroBot-Plugin/plugin/chrev" // 中文字符概述
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/coser" // 三次元小姐姐"github.com/FloatTech/ZeroBot-Plugin/plugin/coser" // 三次元小姐姐"github.com/FloatTech/ZeroBot-Plugin/plugin/coser" // 三次元小姐姐"github.com/FloatTech/ZeroBot-Plugin/plugin/coser"            // 三次元小姐姐
	_“github.com/FloatTech/ZeroBot-Plugin/plugin/cpstory”"github.com/FloatTech/ZeroBot-Plugin/plugin/cpstory" 
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/dailynews" // 今日早报"github.com/FloatTech/ZeroBot-Plugin/plugin/dailynews" // 今日早报"github.com/FloatTech/ZeroBot-Plugin/plugin/dailynews" // 今日早报"github.com/FloatTech/ZeroBot-Plugin/plugin/dailynews"        // 今日早报
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/danbooru" // DeepDanbooru二次元图标签识别"github.com/FloatTech/ZeroBot-Plugin/plugin/danbooru" // DeepDanbooru二次元图标签识别"github.com/FloatTech/ZeroBot-Plugin/plugin/danbooru" // DeepDanbooru二次元图标签识别"github.com/FloatTech/ZeroBot-Plugin/plugin/danbooru"         // DeepDanbooru二次元图标签识别
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/diana" // 嘉心糖行动"github.com/FloatTech/ZeroBot-Plugin/plugin/diana" // 嘉心糖行动"github.com/FloatTech/ZeroBot-Plugin/plugin/diana" // 嘉心糖行动"github.com/FloatTech/ZeroBot-Plugin/plugin/diana"            // 嘉心糖发病
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/dish" // 程序员烹饪指南"github.com/FloatTech/ZeroBot-Plugin/plugin/dish" // 程序员烹饪指南"github.com/FloatTech/ZeroBot-Plugin/plugin/dish" // 程序员烹饪指南"github.com/FloatTech/ZeroBot-Plugin/plugin/dish"             // 程序员做饭指南
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/drawlots" // 多功能抽签 "github.com/FloatTech/ZeroBot-Plugin/plugin/drawlots" // 多功能抽签"github.com/FloatTech/ZeroBot-Plugin/plugin/drawlots" // 多功能抽签"github.com/FloatTech/ZeroBot-Plugin/plugin/drawlots"         // 多功能抽签
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/dress" // 女装"github.com/FloatTech/ZeroBot-Plugin/plugin/dress" // 女装"github.com/FloatTech/ZeroBot-Plugin/plugin/dress" // 女装"github.com/FloatTech/ZeroBot-Plugin/plugin/dress"            // 女装
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/driftbottle" // 漂流瓶"github.com/FloatTech/ZeroBot-Plugin/plugin/driftbottle" // 漂流瓶"github.com/FloatTech/ZeroBot-Plugin/plugin/driftbottle" // 漂流瓶"github.com/FloatTech/ZeroBot-Plugin/plugin/driftbottle"      // 漂流瓶
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/emojimix" // 合成表情符号"github.com/FloatTech/ZeroBot-Plugin/plugin/emojimix"         // 合成emoji
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/event" // 好友申请群聊邀请事件处理"github.com/FloatTech/ZeroBot-Plugin/plugin/event"            // 好友申请群聊邀请事件处理
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/font" // 渲染任意文字到图片"github.com/FloatTech/ZeroBot-Plugin/plugin/font"             // 渲染任意文字到图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/fortune" // 运势"github.com/FloatTech/ZeroBot-Plugin/plugin/fortune"          // 运势
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/funny" // 笑话"github.com/FloatTech/ZeroBot-Plugin/plugin/funny"            // 笑话
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/genshin" // 原神抽卡"github.com/FloatTech/ZeroBot-Plugin/plugin/genshin"          // 原神抽卡
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/gif" // 制图"github.com/FloatTech/ZeroBot-Plugin/plugin/gif"              // 制图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/github" // 搜索 GitHub 仓库"github.com/FloatTech/ZeroBot-Plugin/plugin/github"           // 搜索GitHub仓库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/guessmusic" // 猜歌"github.com/FloatTech/ZeroBot-Plugin/plugin/guessmusic"       // 猜歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/heisi" // 黑丝"github.com/FloatTech/ZeroBot-Plugin/plugin/heisi"            // 黑丝
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hitokoto" // 一言"github.com/FloatTech/ZeroBot-Plugin/plugin/hitokoto"         // 一言
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hs" // 炉石"github.com/FloatTech/ZeroBot-Plugin/plugin/hs"               // 炉石
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hyaku" // 百人一个"github.com/FloatTech/ZeroBot-Plugin/plugin/hyaku"            // 百人一首
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/imgfinder" // 关键字搜图"github.com/FloatTech/ZeroBot-Plugin/plugin/imgfinder"        // 关键字搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/inject" // 注入指令"github.com/FloatTech/ZeroBot-Plugin/plugin/inject"           // 注入指令
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jandan" // 煎蛋网无聊图"github.com/FloatTech/ZeroBot-Plugin/plugin/jandan"           // 煎蛋网无聊图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jiami" // 兽语加密"github.com/FloatTech/ZeroBot-Plugin/plugin/jiami"            // 兽语加密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jptingroom" // 日语听力学习材料"github.com/FloatTech/ZeroBot-Plugin/plugin/jptingroom"       // 日语听力学习材料
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/juejuezi"         // 绝绝子生成器
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/kfccrazythursday"  // 疯狂星期四
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/lolicon"           // 萝莉控随机图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/magicprompt"       // magicprompt 吟唱提示
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/mcfish"            // 钓鱼模拟器
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/midicreate"        // 简易midi音乐制作
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moegoe"           // 日韩 VITS 模型拟声
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyu"             // 摸鱼
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyucalendar"     // 摸鱼人日历
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/music"            // 点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nativesetu"       // 本地涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nbnhhsh"          // 拼音首字母缩写释义工具
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nihongo"          // 日语语法学习
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/novel"            // 铅笔小说网搜索
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nsfw"             // nsfw图片识别
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nwife"            // 本地老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/omikuji"          // 浅草寺求签
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/qqwife"           // 一群一天一夫一妻制群老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/quan"             // QQ权重查询
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/qzone"            // qq空间表白墙
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/realcugan"        // realcugan清晰术
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/reborn"           // 投胎
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/runcode"          // 在线运行代码
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/saucenao"         // 以图搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/scale"            // 叔叔的AI二次元图片放大
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/score"            // 分数
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/setutime"         // 来份涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shadiao"          // 沙雕app
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shindan"          // 测定
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/steam"            // steam相关
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tarot"            // 抽塔罗牌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tiangou"          // 舔狗日记
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tracemoe"         // 搜番
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/translation"      // 翻译
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vitsnyaru"        // vits猫雷
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vtbmusic"         // vtb点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vtbquotation"     // vtb语录
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wallet"           // 钱包
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wangyiyun"        // 网易云音乐热评
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wantquotes"       // 据意查句
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/warframeapi"      // warframeAPI插件
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wenben"           // 文本指令大全
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wenxinvilg"       // 百度文心AI画图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wife"             // 抽老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wordcount"        // 聊天热词
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wordle"           // 猜单词
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ygo"              // 游戏王相关插件
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ymgal"            // 月幕galgame

	// _ "github.com/FloatTech/ZeroBot-Plugin/plugin/wtf"           // 鬼东西

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	//                      ^^^^^^^中优先级区^^^^^^^                      //
	//               ^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^^^               //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------中优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// ----------------------------低优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvvv //
	//               vvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvv               //
	//                      vvvvvvv低优先级区vvvvvvv                      //
	//                          vvvvvvvvvvvvvv                          //
	//                               vvvv                               //

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/curse" // 骂人

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ai_reply" // 人工智能回复

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/thesaurus" // 词典匹配回复

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/breakrepeat" // 打断复读

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	//                      ^^^^^^^低优先级区^^^^^^^                      //
	//               ^^^^^^^^^^^^^^低优先级区^^^^^^^^^^^^^^               //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^低优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------低优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// -----------------------以下为内置依赖，勿动------------------------ //
	“github.com/FloatTech/floatbox/process”
	“github.com/sirupsen/logrus”
	零“github.com/wdvxdr1123/ZeroBot”
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"

	// webctrl "github.com/FloatTech/zbputils/control/web"

	"github.com/FloatTech/ZeroBot-Plugin/kanban/banner"
	// -----------------------以上为内置依赖，勿动------------------------ //
)

type zbpcfg struct {
	Z zero.Config        `json:"zero"`
	W []*driver.WSClient `json:"ws"`
	S []*driver.WSServer `json:"wss"`
}

var config zbpcfg

func init() {
	sus := make([]int64, 0, 16)
	// 解析命令行参数
	d := flag.Bool("d", false, "Enable debug level log and higher.")
	w := flag.Bool("w", false, "Enable warning level log and higher.")
	h := flag.Bool("h", false, "Display this help.")
	// g := flag.String("g", "127.0.0.1:3000", "Set webui url.")
	// 直接写死 AccessToken 时，请更改下面第二个参数
	token := flag.String("t", "", "Set AccessToken of WSClient.")
	// 直接写死 URL 时，请更改下面第二个参数
	url := flag.String("u", "ws://127.0.0.1:6700", "Set Url of WSClient.")
	// 默认昵称
	adana := flag.String("n", "测试", "Set default nickname.")
	prefix := flag.String("p", "/", "Set command prefix.")
	runcfg := flag.String("c", "", "Run from config file.")
	save := flag.String("s", "", "Save default config to file and exit.")
	late := flag.Uint("l", 233, "Response latency (ms).")
	rsz := flag.Uint("r", 4096, "Receiving buffer ring size.")
	maxpt := flag.Uint("x", 4, "Max process time (min).")
	markmsg := flag.Bool("m", false, "Don't mark message as read automatically")

	flag.Parse()

	if *h {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *d && !*w {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if *w {
		logrus.SetLevel(logrus.WarnLevel)
	}

	for _, s := range flag.Args() {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		sus = append(sus, i)
	}

	// 通过代码写死的方式添加主人账号
	 sus = append(sus, 2662596454)
	// sus = append(sus, 87654321)

	// 启用 webui
	// go webctrl.RunGui(*g)

	if *runcfg != "" {
		f, err := os.Open(*runcfg)
		if err != nil {
			panic(err)
		}
		config.W = make([]*driver.WSClient, 0, 2)
		err = json.NewDecoder(f).Decode(&config)
		f.Close()
		if err != nil {
			panic(err)
		}
		config.Z.Driver = make([]zero.Driver, len(config.W)+len(config.S))
		for i, w := range config.W {
			config.Z.Driver[i] = w
		}
		for i, s := range config.S {
			config.Z.Driver[i+len(config.W)] = s
		}
		logrus.Infoln("[main] 从", *runcfg, "读取配置文件")
		return
	}
	config.W = []*driver.WSClient{driver.NewWebSocketClient(*url, *token)}
	config.Z = zero.Config{
		NickName:       append([]string{*adana}, "ATRI", "atri", "亚托莉", "アトリ"),
		CommandPrefix:  *prefix,
		SuperUsers:     sus,
		RingLen:        *rsz,
		Latency:        time.Duration(*late) * time.Millisecond,
		MaxProcessTime: time.Duration(*maxpt) * time.Minute,
		MarkMessage:    !*markmsg,
		Driver:         []zero.Driver{config.W[0]},
	}

	if *save != "" {
		f, err := os.Create(*save)
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(f).Encode(&config)
		f.Close()
		if err != nil {
			panic(err)
		}
		logrus.Infoln("[main] 配置文件已保存到", *save)
		os.Exit(0)
	}
}

func main() {
	if !strings.Contains(runtime.Version(), "go1.2") { // go1.20之前版本需要全局 seed，其他插件无需再 seed
		rand.Seed(time.Now().UnixNano()) //nolint: staticcheck
	}
	// 帮助
	zero.OnFullMatchGroup([]string{"help", "/help", ".help", "菜单"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(banner.Banner, "\n管理发送\"/服务列表\"查看 bot 功能\n发送\"/用法name\"查看功能用法"))
		})
	zero.OnFullMatch("查看zbp公告", zero.OnlyToMe, zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(strings.ReplaceAll(kanban.Kanban(), "\t", "")))
		})
	zero.RunAndBlock(&config.Z, process.GlobalInitMutex.Unlock)
}
