package main

import (
	"fmt"
	"os"
	"strings"

	"go.pedge.io/pkg/cobra"
	"go.pedge.io/protoeasy"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	compilerDirectives := &protoeasy.CompilerDirectives{}

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "Compile protocol buffers files",
		Long:  "Compile protocol buffers files",
		Run: pkgcobra.RunFixedArgs(
			1,
			func(args []string) error {
				return compile(args[0], compilerDirectives)
			},
		),
	}
	bindCompilerDirectives(rootCmd.Flags(), compilerDirectives)

	return rootCmd.Execute()
}

func compile(dirPath string, compilerDirectives *protoeasy.CompilerDirectives) error {
	argsList, err := protoeasy.NewCompiler(
		protoeasy.NewProtoSpecProvider(
			protoeasy.ProtoSpecProviderOptions{},
		),
		protoeasy.CompilerOptions{},
	).ArgsList(
		dirPath,
		compilerDirectives,
	)
	if err != nil {
		return err
	}
	for _, args := range argsList {
		fmt.Println(strings.Join(args, " "))
	}
	return nil
}

func bindCompilerDirectives(flagSet *pflag.FlagSet, compilerDirectives *protoeasy.CompilerDirectives) error {
	flagSet.StringSliceVar(&compilerDirectives.ExcludeFilePatterns, "exclude", []string{}, "Exclude file patterns.")
	flagSet.StringSliceVarP(&compilerDirectives.ProtoPaths, "proto_path", "I", []string{}, "Additional directories for proto imports.")
	flagSet.StringVar(&compilerDirectives.OutDirPath, "out", "", "Customize out directory path.")
	flagSet.BoolVar(&compilerDirectives.Cpp, "cpp", false, "Output cpp files.")
	flagSet.StringVar(&compilerDirectives.CppRelOutDirPath, "cpp_rel_out", "", "Output cpp files.")
	flagSet.BoolVar(&compilerDirectives.Csharp, "csharp", false, "Output csharp files.")
	flagSet.StringVar(&compilerDirectives.CsharpRelOutDirPath, "csharp_rel_out", "", "Output csharp files.")
	flagSet.BoolVar(&compilerDirectives.ObjectiveC, "objectivec", false, "Output objectivec files.")
	flagSet.StringVar(&compilerDirectives.ObjectiveCRelOutDirPath, "objectivec_rel_out", "", "Output objectivec files.")
	flagSet.BoolVar(&compilerDirectives.Python, "python", false, "Output python files.")
	flagSet.StringVar(&compilerDirectives.PythonRelOutDirPath, "python_rel_out", "", "Output python files.")
	flagSet.BoolVar(&compilerDirectives.Ruby, "ruby", false, "Output ruby files.")
	flagSet.StringVar(&compilerDirectives.RubyRelOutDirPath, "ruby_rel_out", "", "Output ruby files.")
	flagSet.BoolVar(&compilerDirectives.Go, "go", false, "Output go files.")
	flagSet.StringVar(&compilerDirectives.GoRelOutDirPath, "go_rel_out", "", "Output go files.")
	flagSet.BoolVar(&compilerDirectives.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&compilerDirectives.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway files.")
	flagSet.BoolVar(&compilerDirectives.Protolog, "protolog", false, "Output protolog files.")
	return nil
}
