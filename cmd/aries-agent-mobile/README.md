# Aries Agent Mobile

Mobile bindings for the [Aries Framework Go](https://github.com/hyperledger/aries-framework-go) library.
> Note: these bindings are experimental and are subject to frequent changes.

## 1. Requirements

- [Golang](https://golang.org/doc/install) >= 1.13
- [Android SDK](https://developer.android.com/studio/install) (via Android Studio)
- [Android NDK](https://developer.android.com/ndk/downloads)
- [Xcode](https://developer.apple.com/xcode/) (macOS only)
- Make
    - [Windows](http://gnuwin32.sourceforge.net/packages/make.htm)
    - [macOS](https://brew.sh/) (via Homebrew)
    - Linux (pre-installed)


## 2. Build

Export (or set) the following variables
- `ANDROID_HOME` - the location of the installed Android SDK
- `ANDROID_NDK_HOME` - the location of the installed Android NDK

### 2.1 Make All

```bash
$ make all
```

### 2.2 Make Bindings

#### a. All bindings
```bash
$ make bindings
```

#### b. Android bindings
```bash
$ make bindings-android
```

#### c. iOS bindings
> Please note that this can only be run on macOS.
```bash
$ make bindings-ios
```

## 3. Usage

### 3.1. Android

#### a. Importing the generated binding as a module in Android Studio
- In the menu of your Android Studio project, go to **File>Project Structure**.
- A modal will be displayed and on the left click on **Modules**.
- In the section title _Modules_ click on the **+**.
- Another modal will be displayed, scroll down and select _Import .JAR/.AAR Package_ and press **Next**.
- In the _File name_ field, enter the path to the `aries-agent.aar` file and click **Finish**.
- Select **Apply** if applicable and then **OK**.
- Reopen the _Project Structure_ modal and on the left click on **Dependencies**.
- Click on **app**. In the section titled _Declared Dependencies_, click on the **+**.
- Click on **Module Dependency** and select `aries-agent.aar`.
- Click **OK** and select **Apply** if applicable and then **OK**.

#### b. Code sample
This is an example of how the imported module can be used:
```java
import org.hyperledger.aries.api.AriesController;
import org.hyperledger.aries.api.IntroduceController;
import org.hyperledger.aries.ariesagent.Ariesagent;
import org.hyperledger.aries.models.RequestEnvelope;
import org.hyperledger.aries.models.ResponseEnvelope;
import org.hyperledger.aries.config.Options;

import java.nio.charset.StandardCharsets;
/*
...
*/
        // create options
        Options opts = new Options();
        opts.setAgentURL("http://example.com");
        opts.setUseLocalAgent(false);

        ResponseEnvelope res = new ResponseEnvelope();
        try {
            // create an aries agent instance
            AriesController a = Ariesagent.new_(opts);

            // create a controller
            IntroduceController i = a.getIntroduceController();

            // perform an operation
            byte[] data = "{}".getBytes(StandardCharsets.UTF_8);
            res = i.actions(new RequestEnvelope(data));
        } catch (Exception e) {
            e.printStackTrace();
        }

        String actionsResponse = new String(res.getPayload(), StandardCharsets.UTF_8);
        System.out.println(actionsResponse);
```


### 3.2. iOS

#### a. Importing the generated binding as a framework in Xcode
- In the menu of your Xcode project, go to **File>Add Files to "your project name"...**.
- In the displayed modal, navigate to the path of your `AriesAgent.framework` file and click **Add**.

#### b. Code sample
This is an example of how the imported framework can be used:
```objc
#import <AriesAgent/Ariesagent.h>
/*
...
*/
    NSError *error = nil;

    // create options
    ConfigOptions *opts = ConfigNew();
    // [opts setAgentURL:@"http://example.com"];
    [opts setUseLocalAgent:true];
    
    // create an aries agent instance
    ApiAriesController *ac = (ApiAriesController*) AriesagentNew(opts, &error);
    if(error) {
        NSLog(@"error creating an aries agent: %@", error);
    }
    
    // create a controller
    ApiVerifiableController *ic = (ApiVerifiableController*) [ac getVerifiableController:&error];
    if(error) {
        NSLog(@"error creating an verifiable controller instance: %@", error);
    }

    // perform an operation
    NSData *data = [@"" dataUsingEncoding:NSUTF8StringEncoding];
    ModelsRequestEnvelope *req = ModelsNewRequestEnvelope(data);
    ModelsResponseEnvelope *resp = [ic getCredentials:req];
    if(resp.error) {
        NSLog(@"error getting credentials: %@", resp.error.message);
    } else {
        NSString *credResp = [[NSString alloc] initWithData:resp.payload encoding:NSUTF8StringEncoding];
        NSLog(@"credentials response: %@", credResp);
    }
```


### 3.3. Demo apps

For examples of mobile apps built with the aries-agent-mobile bindings, see [https://github.com/trustbloc/aries-examples](https://github.com/trustbloc/aries-examples).


## 4. Test

```bash
$ make unit-test
```


## 5. Project Structure

The project structure for the mobile bindings can be found [here](doc/project_structure.md).

## 6. Contribute

See the [guidelines](https://github.com/hyperledger/aries-framework-go/blob/master/.github/CONTRIBUTING.md) from the parent project.
