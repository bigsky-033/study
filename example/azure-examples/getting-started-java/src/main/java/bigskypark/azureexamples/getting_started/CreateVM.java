package bigskypark.azureexamples.getting_started;

import com.azure.core.http.policy.HttpLogDetailLevel;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.Region;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.AzureAuthorityHosts;
import com.azure.identity.EnvironmentCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;
import com.azure.resourcemanager.compute.models.KnownLinuxVirtualMachineImage;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;

import java.nio.file.Files;
import java.nio.file.Paths;

public class CreateVM {
    public static void main(String[] args) throws Exception {
        final var userName = "bigsky033";
        final var sshKey = Files.readString(Paths.get("getting-started-java/keys/id_rsa.pub"));

        final var credential = new EnvironmentCredentialBuilder()
                .authorityHost(AzureAuthorityHosts.AZURE_PUBLIC_CLOUD)
                .build();

        final var profile = new AzureProfile(AzureEnvironment.AZURE);

        final var azureResourceManager = AzureResourceManager.configure()
                .withLogLevel(HttpLogDetailLevel.BASIC)
                .authenticate(credential, profile)
                .withDefaultSubscription();

        final var linuxVM = azureResourceManager.virtualMachines().define("testLinuxVM")
                .withRegion(Region.ASIA_EAST)
                .withNewResourceGroup("sampleVmResourceGroup")
                .withNewPrimaryNetwork("10.0.0.0/24")
                .withPrimaryPrivateIPAddressDynamic()
                .withoutPrimaryPublicIPAddress()
                .withPopularLinuxImage(KnownLinuxVirtualMachineImage.UBUNTU_SERVER_18_04_LTS)
                .withRootUsername(userName)
                .withSsh(sshKey)
                .withUnmanagedDisks()
                .withSize(VirtualMachineSizeTypes.STANDARD_D3_V2)
                .create();
    }
}
