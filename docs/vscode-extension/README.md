# Installing the Cyone VS Code Extension

## Prerequisites

- [Node.js](https://nodejs.org/)
- [npm](https://www.npmjs.com/)
- [Visual Studio Code](https://code.visualstudio.com/)

## Steps to Compile and Install

1. **Install `vsce`**

   Open your terminal and run:

   ```
   npm install -g vsce
   ```

2. **Package the Extension**

   Navigate to the directory with your extension project and run:

   ```
   vsce package
   ```

   This creates a `.vsix` file in the directory.

3. **Install the Extension in VS Code**

   Install the `.vsix` file using the following command:

   ```
   code --install-extension cyone-0.0.1.vsix
   ```

   Replace `cyone-0.0.1.vsix` with the name of your `.vsix` file.

## Reporting Issues

If you encounter any issues, please report them on our [GitHub Issues page](https://github.com/isakruas/cyone-assembly/issues). Provide as much detail as possible to help us address the problem.

## Contributing

We welcome contributions to improve the Cyone extension! If you have suggestions or fixes, please submit a pull request on our [GitHub repository](https://github.com/isakruas/cyone-assembly). Follow these steps to contribute:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your modifications and commit them.
4. Push your changes to your fork.
5. Open a pull request to merge your changes into the main repository.

---

For any additional questions or help, feel free to reach out!
